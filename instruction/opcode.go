package instruction

import (
	"8086-disassembler/util"
	"fmt"
)

type Opcode int

const (
	Mov Opcode = iota + 1
)

func (op Opcode) String() string {
	switch op {
	case Mov:
		return "MOV"
	default:
		return "Unknown"
	}
}

func DecodeOpcode(opcodeSection byte) (Opcode, error) {
	codes := map[byte]Opcode{
		// 100010
		0x22: Mov,
	}

	opcode, ok := codes[opcodeSection]
	if !ok {
		return 0, fmt.Errorf("unknown opcode 0x%X (%s).", opcodeSection, util.RenderByte(opcodeSection))
	}
	return opcode, nil
}
