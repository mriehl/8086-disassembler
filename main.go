package main

import (
	"8086-disassembler/instruction"
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	listings := []string{"asm/37", "asm/38"}

	for _, listing := range listings {
		fmt.Printf("=== Listing %s ===\n", listing)
		handleListing(listing)
	}
}

func handleListing(listing string) {
	file, err := os.Open(listing)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	instruction_reader := bufio.NewReader(file)

	decoded_asm, err := os.Create(fmt.Sprintf("asm_decoded/%s.asm", strings.Split(listing, "/")[1]))
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

	// max inst length is 6 byte
	buf := make([]byte, 6)

	requestInstSize := func(n int) []byte {
		remaining := make([]byte, n-1)
		_, err := instruction_reader.Read(remaining)
		if err != nil {
			panic(err)
		}
		copy(buf[1:], remaining)
		return buf[:n]
	}

	for {
		firstByte, err := instruction_reader.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		buf[0] = firstByte
		opcode, err := instruction.DecodeOpcode(firstByte)
		if err != nil {
			panic(err)
		}
		var currentInst instruction.StringerInstruction
		switch opcode {
		case instruction.Mov:
			currentInst, err = instruction.DecodeMovInstruction(requestInstSize)
		case instruction.MovImmediateToReg:
			panic("immediate mov")
		default:
			panic(fmt.Errorf("unexpected opcode %s", opcode))
		}
		if err != nil {
			fmt.Printf("\tINVALID: %v\n", err)
		} else {
			inst := strings.ToLower(currentInst.AsStringInstruction())
			fmt.Printf("%s\t from %+v\n", inst, currentInst)
			_, err = decoded_asm_writer.WriteString(inst + "\n")
			if err != nil {
				panic(err)
			}
		}
	}
}
