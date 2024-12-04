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

func DecodeOpcode(b byte) (Opcode, error) {
	section := b & 0xFC
	codes := map[byte]Opcode{
		0x88: Mov,
	}

	opcode, ok := codes[section]
	if !ok {
		return 0, fmt.Errorf("unknown opcode 0x%X (%s).", b, util.RenderBytes([]byte{b}))
	}
	return opcode, nil
}
