package instructions

import (
	"8086-disassembler/decoder/fields"
	"8086-disassembler/util"
	"encoding/binary"
	"fmt"
	"slices"
)

type MovImmediateToRegMemInstruction struct {
	Raw          []byte
	InstBuf      string
	Opcode       fields.Opcode
	W            fields.W
	Mod          fields.Mod
	Dest         interface{}
	SourceValue  uint16
	SourceIsWord bool
}

func (mov MovImmediateToRegMemInstruction) AsStringInstruction() string {
	var size string
	if mov.W == fields.Byte {
		size = "byte"
	} else {
		size = "word"
	}
	explicitValue := fmt.Sprintf("%s %d", size, mov.SourceValue)
	return fmt.Sprintf("mov %s, %s", mov.Dest, explicitValue)
}

func DecodeMovImmediateToRegMem(byte1 byte, opcode fields.Opcode, requestFurtherBytes func(int) []byte) (util.InstructionType, error) {
	// | ________ | __  000 ___ | byte    | byte    | byte | byte      |
	// | 1100011w | mod     r/m | disp-lo | disp-hi | data | data(w=1) |
	w, err := fields.DecodeW(byte1 & 0x1)
	if err != nil {
		return nil, err
	}
	var byte2 = requestFurtherBytes(1)[0]
	mod, err := fields.DecodeMod(byte2 >> 6)
	if err != nil {
		return nil, err
	}

	// TODO duplicated rm + disp fuckery
	rawInstruction := []byte{byte1, byte2}
	var rm interface{}
	rmSection := byte2 & 0x7
	if mod == fields.RegisterMode {
		rm, err = fields.DecodeReg(rmSection, w)
		if err != nil {
			return nil, err
		}
	} else {
		additionalRequired := fields.DecodeTrailingMemoryLength(rmSection, mod)
		additionalBytes := requestFurtherBytes(int(additionalRequired))
		rawInstruction = slices.Concat(rawInstruction, additionalBytes)
		if additionalRequired > 0 {
			copy(rawInstruction[2:], additionalBytes)
		}
		rm, err = fields.DecodeMemoryAddress(rmSection, mod, additionalBytes)
		if err != nil {
			panic(err)
		}
	}

	var immediateValue uint16
	if w == fields.Byte {
		data := requestFurtherBytes(1)
		rawInstruction = slices.Concat(rawInstruction, data)
		immediateValue = uint16(data[0])
	} else {
		data := requestFurtherBytes(2)
		rawInstruction = slices.Concat(rawInstruction, data)
		immediateValue = binary.LittleEndian.Uint16(data)
	}

	return MovImmediateToRegMemInstruction{
		Raw:          rawInstruction,
		InstBuf:      util.RenderBytes(rawInstruction),
		Opcode:       opcode,
		Mod:          mod,
		W:            w,
		Dest:         rm,
		SourceValue:  immediateValue,
		SourceIsWord: w == fields.Word,
	}, nil
}

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

func DecodeMovImmediateToReg(byte1 byte, opcode fields.Opcode, requestFurtherBytes func(int) []byte) (util.InstructionType, error) {
	// | ____ _ ___ | data | data |
	// | 1011 w reg |      |      |
	w, err := fields.DecodeW(byte1 >> 3 & 0x1)
	if err != nil {
		return nil, err
	}
	var immediate []byte
	var immediateValue uint16
	if w == fields.Byte {
		immediate = requestFurtherBytes(1)
		immediateValue = uint16(immediate[0])
	} else {
		immediate = requestFurtherBytes(2)
		immediateValue = binary.LittleEndian.Uint16(immediate)
	}

	reg, err := fields.DecodeReg(byte1&0x7, w)

	raw := slices.Concat([]byte{byte1}, immediate)

	return MovImmediateToRegInstruction{
		Raw:         raw,
		InstBuf:     util.RenderBytes(raw),
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

func DecodeMovRmToFromReg(byte1 byte, opcode fields.Opcode, requestFurtherBytes func(int) []byte) (util.InstructionType, error) {
	// | ______ _ _ | __  ___ ___ |
	// | 100010 d w | mod reg r/m |
	byte2 := requestFurtherBytes(1)[0]

	mod, err := fields.DecodeMod(byte2 >> 6 & 0x3)
	if err != nil {
		return nil, err
	}
	w, err := fields.DecodeW(byte1 & 0x1)
	if err != nil {
		return nil, err
	}
	d, err := fields.DecodeD(byte1 >> 1 & 0x1)
	if err != nil {
		return nil, err
	}

	reg, err := fields.DecodeReg(byte2>>3&0x7, w)
	if err != nil {
		return nil, err
	}

	rawInstruction := make([]byte, 2, 4)
	rawInstruction[0] = byte1
	rawInstruction[1] = byte2

	var rm interface{}
	rmSection := byte2 & 0x7
	if mod == fields.RegisterMode {
		rm, err = fields.DecodeReg(rmSection, w)
		if err != nil {
			return nil, err
		}
	} else {
		additionalRequired := fields.DecodeTrailingMemoryLength(rmSection, mod)
		additionalBytes := requestFurtherBytes(int(additionalRequired))
		rawInstruction = rawInstruction[:additionalRequired+2]
		if additionalRequired > 0 {
			copy(rawInstruction[2:], additionalBytes)
		}
		rm, err = fields.DecodeMemoryAddress(rmSection, mod, additionalBytes)
		if err != nil {
			panic(err)
		}
	}

	inst := MovRmToFromRegInstruction{
		Raw:     rawInstruction,
		InstBuf: util.RenderBytes(rawInstruction),
		Opcode:  opcode,
		Mod:     mod,
		W:       w,
		D:       d,
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
