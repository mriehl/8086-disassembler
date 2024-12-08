package instructions

import (
	"8086-disassembler/decoder/fields"
	"8086-disassembler/util"
	"bufio"
	"fmt"
)

type MovInstruction struct {
	Raw     []byte
	InstBuf string
	Opcode  fields.Opcode
	Mod     fields.Mod
	W       fields.W
	D       fields.D
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

func DecodeMovInstruction(opcode fields.Opcode, requestInstSize func(int) []byte) (util.InstructionType, error) {
	// | ______ _ _ | __  ___ ___ |
	// | opcode d w | mod reg r/m |
	raw := requestInstSize(2)

	bits := util.RenderBytes(raw)

	mod, err := fields.DecodeMod(raw[1] >> 6 & 0x3)
	if err != nil {
		return nil, err
	}
	w, err := fields.DecodeW(raw[0] & 0x1)
	if err != nil {
		return nil, err
	}
	d, err := fields.DecodeD(raw[0] >> 1 & 0x1)
	if err != nil {
		return nil, err
	}

	reg, err := fields.DecodeReg(raw[1]>>3&0x7, w)
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
	if mod == fields.RegisterToRegister {
		rm, err = fields.DecodeReg(raw[1]&0x7, w)
		if err != nil {
			return nil, err
		}
	} else {
		panic("rm is memory address")
	}

	if d == fields.RegIsDest {
		inst.Dest = reg
		inst.Source = rm
	} else {
		inst.Source = reg
		inst.Dest = rm
	}

	return inst, nil
}
