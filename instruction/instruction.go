package instruction

import (
	"8086-disassembler/util"
	"fmt"
)

type MovInstruction struct {
	Raw    []byte
	Bits   string
	Opcode Opcode
	Mod    Mod
	W      W
	D      D
	Source interface{}
	Dest   interface{}
}

func (mov *MovInstruction) AsStringInst() string {
	return fmt.Sprintf("mov %s, %s", mov.Dest, mov.Source)
}

func DecodeMovInstruction(raw []byte) (*MovInstruction, error) {
	bits := util.RenderBytes(raw)
	if len(raw) != 2 {
		return nil, fmt.Errorf("expected 2 bytes for instruction but got %s", bits)
	}

	// | ______ _ _ | __  ___ ___ |
	// | opcode d w | mod reg r/m |

	opcode, err := DecodeOpcode(raw[0] >> 2 & 0x3f)
	if err != nil {
		return nil, err
	}
	mod, err := DecodeMod(raw[1] >> 6 & 0x3)
	if err != nil {
		return nil, err
	}
	w, err := DecodeW(raw[0] & 0x1)
	if err != nil {
		return nil, err
	}
	d, err := DecodeD(raw[0] >> 1 & 0x1)
	if err != nil {
		return nil, err
	}

	reg, err := DecodeReg(raw[1]>>3&0x7, w)
	if err != nil {
		return nil, err
	}

	inst := &MovInstruction{
		Raw:    raw,
		Bits:   bits,
		Opcode: opcode,
		Mod:    mod,
		W:      w,
		D:      d,
	}

	var rm interface{}
	if mod == RegisterToRegister {
		rm, err = DecodeReg(raw[1]&0x7, w)
		if err != nil {
			return nil, err
		}
	} else {
		panic("rm is memory address")
	}

	if d == RegIsDest {
		inst.Dest = reg
		inst.Source = rm
	} else {
		inst.Source = reg
		inst.Dest = rm
	}

	return inst, nil
}
