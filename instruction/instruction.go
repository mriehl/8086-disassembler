package instruction

import (
	"8086-disassembler/util"
	"bufio"
	"fmt"
)

type StringerInstruction interface {
	AsStringInstruction() string
}

type MovInstruction struct {
	Raw     []byte
	InstBuf string
	Opcode  Opcode
	Mod     Mod
	W       W
	D       D
	Source  interface{}
	Dest    interface{}
}

func (mov MovInstruction) AsStringInstruction() string {
	return fmt.Sprintf("mov %s, %s", mov.Dest, mov.Source)
}

func readSubsequentBytes(n int, buf []byte, reader *bufio.Reader) error {
	remaining := make([]byte, n)
	_, err := reader.Read(remaining)
	if err != nil {
		return err
	}
	copy(buf[1:], remaining)
	return nil
}

func DecodeMovInstruction(requestInstSize func(int) []byte) (StringerInstruction, error) {
	// | ______ _ _ | __  ___ ___ |
	// | opcode d w | mod reg r/m |
	raw := requestInstSize(2)

	bits := util.RenderBytes(raw)

	opcode, err := DecodeOpcode(raw[0])
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

	inst := MovInstruction{
		Raw:     raw,
		InstBuf: bits,
		Opcode:  opcode,
		Mod:     mod,
		W:       w,
		D:       d,
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
