package instruction

import (
	"8086-disassembler/util"
	"fmt"
)

type Instruction struct {
	Raw    []byte
	Bits   string
	Opcode Opcode
	Mod    Mod
	W      W
	D      D
	Reg    Reg
}

func DecodeInstruction(raw []byte) (*Instruction, error) {
	bits := util.RenderBytes(raw)
	if len(raw) != 2 {
		return nil, fmt.Errorf("expected 2 bytes for instruction but got %s", bits)
	}

	// | ______ _ _ | __  ___ ___ |
	// | opcode d w | mod reg r/m |

	opcode, err := DecodeOpcode(raw[0])
	if err != nil {
		return nil, err
	}
	mod, err := DecodeMod(raw[1])
	if err != nil {
		return nil, err
	}
	w, err := DecodeW(raw[1])
	if err != nil {
		return nil, err
	}
	d, err := DecodeD(raw[1])
	if err != nil {
		return nil, err
	}

	reg, err := DecodeReg(raw[1], w)
	if err != nil {
		return nil, err
	}

	return &Instruction{
		Raw:    raw,
		Bits:   bits,
		Opcode: opcode,
		Mod:    mod,
		W:      w,
		D:      d,
		Reg:    reg,
	}, nil
}
