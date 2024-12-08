package instructions

import (
	"8086-disassembler/decoder/fields"
	"8086-disassembler/util"
	"encoding/binary"
	"fmt"
)

type MovImmediateToRegInstruction struct {
	Raw         []byte
	InstBuf     string
	Opcode      fields.Opcode
	W           fields.W
	SourceValue uint16
	Dest        fields.Reg
}

func (mov MovImmediateToRegInstruction) AsStringInstruction() string {
	return fmt.Sprintf("mov %s, %d", mov.Dest, mov.SourceValue)
}

func DecodeMovImmediateToReg(firstByte byte, opcode fields.Opcode, requestInstSize func(int) []byte) (util.InstructionType, error) {
	// | ____ _ ___ | data | data |
	// | 1011 w reg |      |      |
	w, err := fields.DecodeW(firstByte >> 3 & 0x1)
	if err != nil {
		return nil, err
	}
	var buf []byte
	var immediateValue uint16
	if w == fields.Byte {
		buf = requestInstSize(2)
		immediateValue = uint16(buf[1])
	} else {
		buf = requestInstSize(3)
		immediateValue = binary.LittleEndian.Uint16(buf[1:3])
	}

	reg, err := fields.DecodeReg(firstByte&0x7, w)

	return MovImmediateToRegInstruction{
		Raw:         buf,
		InstBuf:     util.RenderBytes(buf),
		Opcode:      opcode,
		W:           w,
		SourceValue: immediateValue,
		Dest:        reg,
	}, nil
}

type MovRmToFromRegInstruction struct {
	Raw     []byte
	InstBuf string
	Opcode  fields.Opcode
	Mod     fields.Mod
	W       fields.W
	D       fields.D
	Source  interface{}
	Dest    interface{}
}

func (mov MovRmToFromRegInstruction) AsStringInstruction() string {
	return fmt.Sprintf("mov %s, %s", mov.Dest, mov.Source)
}

func DecodeMovRmToFromReg(firstByte byte, opcode fields.Opcode, requestInstSize func(int) []byte) (util.InstructionType, error) {
	// | ______ _ _ | __  ___ ___ |
	// | 100010 d w | mod reg r/m |
	buf := requestInstSize(2)

	mod, err := fields.DecodeMod(buf[1] >> 6 & 0x3)
	if err != nil {
		return nil, err
	}
	w, err := fields.DecodeW(buf[0] & 0x1)
	if err != nil {
		return nil, err
	}
	d, err := fields.DecodeD(buf[0] >> 1 & 0x1)
	if err != nil {
		return nil, err
	}

	reg, err := fields.DecodeReg(buf[1]>>3&0x7, w)
	if err != nil {
		return nil, err
	}

	inst := MovRmToFromRegInstruction{
		Raw:     buf,
		InstBuf: util.RenderBytes(buf),
		Opcode:  opcode,
		Mod:     mod,
		W:       w,
		D:       d,
	}

	var rm interface{}
	if mod == fields.RegisterToRegister {
		rm, err = fields.DecodeReg(buf[1]&0x7, w)
		if err != nil {
			return nil, err
		}
	} else {
		panic(fmt.Errorf("rm is unimplemented memory address in %s", util.RenderBytes(buf)))
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
