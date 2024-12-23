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
	MovMemToAcc
	MovAccToMem
)

func (op Opcode) String() string {
	switch op {
	case MovRmToFromReg:
		return "MovRmToFromReg"
	case MovImmediateToReg:
		return "MovImmediateToReg"
	case MovImmediateToRegMem:
		return "MovImmediateToRegMem"
	case MovMemToAcc:
		return "MovMemToAcc"
	case MovAccToMem:
		return "MovAccToMem"

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
		// 1010000
	case firstByte>>1 == 0x50:
		return MovMemToAcc, nil
		// 1010001
	case firstByte>>1 == 0x51:
		return MovAccToMem, nil
	}

	return 0, fmt.Errorf("unknown opcode 0x%X (%s).", firstByte, util.RenderByte(firstByte))
}
