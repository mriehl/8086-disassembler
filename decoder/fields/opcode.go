package fields

import (
	"8086-disassembler/util"
	"fmt"
)

type Opcode int

const (
	MovRmToFromReg Opcode = iota + 1
	MovImmediateToReg
	MovImmediateToRegMem
)

func (op Opcode) String() string {
	switch op {
	case MovRmToFromReg:
		return "MovRmToFromReg"
	case MovImmediateToReg:
		return "MovImmediateToReg"
	case MovImmediateToRegMem:
		return "MovImmediateToRegMem"
	default:
		return "Unknown"
	}
}

func DecodeOpcode(firstByte byte) (Opcode, error) {
	switch {
	// === MOVs ===
	// 100010__
	case firstByte>>2 == 0x22:
		return MovRmToFromReg, nil
	// 1011____
	case firstByte>>4 == 0xb:
		return MovImmediateToReg, nil
		// 1100011_
	case firstByte>>1 == 0x63:
		return MovImmediateToRegMem, nil
	}

	return 0, fmt.Errorf("unknown opcode 0x%X (%s).", firstByte, util.RenderByte(firstByte))
}
