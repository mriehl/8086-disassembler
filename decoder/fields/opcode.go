package fields

import (
	"8086-disassembler/util"
	"fmt"
)

type Opcode int

const (
	MovRmToFromReg Opcode = iota + 1
	MovImmediateToReg
)

func (op Opcode) String() string {
	switch op {
	case MovRmToFromReg:
		return "MovRmToFromReg"
	case MovImmediateToReg:
		return "MovImmediateToReg"
	default:
		return "Unknown"
	}
}

func DecodeOpcode(firstByte byte) (Opcode, error) {
	switch {
	// 100010__
	case firstByte>>2 == 0x22:
		return MovRmToFromReg, nil
	// 1011____
	case firstByte>>4 == 0xb:
		return MovImmediateToReg, nil
	}

	return 0, fmt.Errorf("unknown opcode 0x%X (%s).", firstByte, util.RenderByte(firstByte))
}
