package decoder

import (
	"8086-disassembler/decoder/fields"
	"8086-disassembler/decoder/instructions"
	"8086-disassembler/util"
	"bufio"
	"fmt"
	"io"
)

type DecodeResult struct {
	Value util.InstructionType
	Error error
}

func DecodeInstructions(instruction_reader *bufio.Reader) <-chan DecodeResult {
	ch := make(chan DecodeResult)
	go func() {
		requestFurtherBytes := func(n int) []byte {
			additional := make([]byte, n)
			_, err := instruction_reader.Read(additional)
			if err != nil {
				panic(err)
			}
			return additional
		}

		for {
			firstByte, err := instruction_reader.ReadByte()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			opcode, err := fields.DecodeOpcode(firstByte)
			if err != nil {
				panic(err)
			}
			var currentInst util.InstructionType
			switch opcode {
			case fields.MovRmToFromReg:
				currentInst, err = instructions.DecodeMovRmToFromReg(firstByte, opcode, requestFurtherBytes)
				break
			case fields.MovImmediateToReg:
				currentInst, err = instructions.DecodeMovImmediateToReg(firstByte, opcode, requestFurtherBytes)
				break
			case fields.MovImmediateToRegMem:
				currentInst, err = instructions.DecodeMovImmediateToRegMem(firstByte, opcode, requestFurtherBytes)
				break
			default:
				panic(fmt.Errorf("unexpected opcode %s", opcode))
			}
			ch <- DecodeResult{
				Value: currentInst,
				Error: err,
			}
		}

		close(ch)
	}()
	return ch
}
