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
			opcode, err := fields.DecodeOpcode(firstByte)
			if err != nil {
				panic(err)
			}
			var currentInst util.InstructionType
			switch opcode {
			case fields.MovRmToFromReg:
				currentInst, err = instructions.DecodeMovRmToFromReg(firstByte, opcode, requestInstSize)
				break
			case fields.MovImmediateToReg:
				currentInst, err = instructions.DecodeMovImmediateToReg(firstByte, opcode, requestInstSize)
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
