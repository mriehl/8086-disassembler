package main

import (
	"8086-disassembler/instruction"
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("asm/38")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, 2)
	instruction_reader := bufio.NewReader(file)

	for {
		_, err := instruction_reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		instruction, err := instruction.DecodeMovInstruction(buf)
		if err != nil {
			fmt.Printf("\tINVALID: %v\n", err)
		} else {
			fmt.Printf("\t%+v\n", instruction)
		}
	}
}
