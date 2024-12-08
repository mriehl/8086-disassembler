package instruction

import (
	"8086-disassembler/util"
	"fmt"
)

type Opcode int

const (
	Mov Opcode = iota + 1
	MovImmediateToReg
)

func (op Opcode) String() string {
	switch op {
	case Mov:
		return "MOV"
	case MovImmediateToReg:
		return "MOV"
	default:
		return "Unknown"
	}
}

func DecodeOpcode(firstByte byte) (Opcode, error) {
	switch {
	// 100010__
	case firstByte>>2 == 0x22:
		return Mov, nil
	// 1011____
	case firstByte>>4 == 0xb:
		return MovImmediateToReg, nil
	}

	return 0, fmt.Errorf("unknown opcode 0x%X (%s).", firstByte, util.RenderByte(firstByte))
}
