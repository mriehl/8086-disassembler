package main

import (
	"8086-disassembler/decoder"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	listings, err := filepath.Glob("asm/*.asm")
	if err != nil {
		panic(err)
	}

	for _, listing := range listings {
		fmt.Printf("=== Listing %s ===\n", listing)
		handleListing(listing)
	}
}

func handleListing(listing string) {
	file, err := os.Open(strings.Replace(listing, ".asm", "", -1))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	instruction_reader := bufio.NewReader(file)

	decoded_asm, err := os.Create(fmt.Sprintf("asm_decoded/%s", strings.Split(listing, "/")[1]))
	if err != nil {
		panic(err)
	}
	defer decoded_asm.Close()
	decoded_asm_writer := bufio.NewWriter(decoded_asm)
	defer decoded_asm_writer.Flush()
	_, err = decoded_asm_writer.WriteString("bits 16\n\n")
	if err != nil {
		panic(err)
	}

	for result := range decoder.DecodeInstructions(instruction_reader) {

		if result.Error != nil {
			fmt.Printf("\tINVALID: %v\n", err)
		} else {
			inst := strings.ToLower(result.Value.AsStringInstruction())
			fmt.Printf("%s\t from %+v\n", inst, result.Value)
			_, err = decoded_asm_writer.WriteString(inst + "\n")
			if err != nil {
				panic(err)
			}
		}

	}
}
