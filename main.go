package main

import (
	"8086-disassembler/instruction"
	"fmt"
	"io"
	"os"
)

func main() {
	instruction_file, err := os.Open("asm/37")
	if err != nil {
		panic(err)
	}
	defer instruction_file.Close()

	buf := make([]byte, 2)

	for {
		_, err := instruction_file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		instruction, err := instruction.DecodeInstruction(buf)
		if err != nil {
			fmt.Printf("\tINVALID: %v\n", err)
		} else {
			fmt.Printf("\t%+v\n", instruction)
		}
	}
}
