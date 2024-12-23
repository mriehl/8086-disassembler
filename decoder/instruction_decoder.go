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
		var instBytes []byte

		requestFurtherBytes := func(n int) []byte {
			additional := make([]byte, n)
			_, err := instruction_reader.Read(additional)
			instBytes = append(instBytes, additional...)
			if err != nil {
				panic(err)
			}
			return additional
		}

		for {
			instBytes = make([]byte, 1, 6)
			firstByte, err := instruction_reader.ReadByte()
			instBytes[0] = firstByte
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			opcode, err := fields.DecodeOpcode(firstByte)
			if err != nil {
				// TODO should not panic here
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
			case fields.MovAccToMem:
				fallthrough
			case fields.MovMemToAcc:
				currentInst, err = instructions.DecodeMovAccMem(firstByte, opcode, requestFurtherBytes)
				break
			default:
				panic(fmt.Errorf("unexpected opcode %s", opcode))
			}

			if err != nil {
				err = fmt.Errorf("error while decoding instruction (read=%s) due to %v", util.RenderBytes(instBytes), err)
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
