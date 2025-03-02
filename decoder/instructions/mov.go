package instructions

import (
	"8086-disassembler/decoder/fields"
	"8086-disassembler/decoder/types"
	"8086-disassembler/util"
	"encoding/binary"
	"fmt"
)

type MovImmediateToRegMemInstruction struct {
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

	eac, err := EAC(byte2&0x7, mod, w, requestFurtherBytes)
	if err != nil {
		return nil, fmt.Errorf("cannot calculate EAC for %s: %w", opcode, err)
	}

	var immediateValue uint16
	if w == fields.Byte {
		data := requestFurtherBytes(1)
		immediateValue = uint16(data[0])
	} else {
		data := requestFurtherBytes(2)
		immediateValue = binary.LittleEndian.Uint16(data)
	}

	return MovImmediateToRegMemInstruction{
		Opcode:       opcode,
		Mod:          mod,
		W:            w,
		Dest:         eac.EffectiveAddress,
		SourceValue:  immediateValue,
		SourceIsWord: w == fields.Word,
	}, nil
}

type MovImmediateToRegInstruction struct {
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

	return MovImmediateToRegInstruction{
		Opcode:      opcode,
		W:           w,
		SourceValue: immediateValue,
		Dest:        reg,
	}, nil
}

type MovRmToFromRegInstruction struct {
	Opcode fields.Opcode
	Mod    fields.Mod
	W      fields.W
	D      fields.D
	Source interface{}
	Dest   interface{}
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

	eac, err := EAC(byte2&0x7, mod, w, requestFurtherBytes)
	if err != nil {
		return nil, fmt.Errorf("cannot calculate EAC for %s: %w", opcode, err)
	}

	inst := MovRmToFromRegInstruction{
		Opcode: opcode,
		Mod:    mod,
		W:      w,
		D:      d,
	}

	if d == fields.RegIsDest {
		inst.Dest = reg
		inst.Source = eac.EffectiveAddress
	} else {
		inst.Source = reg
		inst.Dest = eac.EffectiveAddress
	}

	return inst, nil
}

type MovAccMem struct {
	Opcode fields.Opcode
	W      fields.W
	Source interface{}
	Dest   interface{}
}

func (mov MovAccMem) AsStringInstruction() string {
	return fmt.Sprintf("mov %s, %s", mov.Dest, mov.Source)
}

func DecodeMovAccMem(byte1 byte, opcode fields.Opcode, requestFurtherBytes func(int) []byte) (util.InstructionType, error) {
	// | ______ _   _ | addr-lo | addr-hi |
	// | 101000 0/1 w | addr-lo | addr-hi |

	w, err := fields.DecodeW(byte1 & 0x1)
	if err != nil {
		return nil, err
	}
	acc, err := fields.DecodeAcc(w)
	if err != nil {
		return nil, err
	}

	var addr int
	switch w {
	case fields.Byte:
		data := requestFurtherBytes(1)
		addr = int(uint8(data[0]))
	case fields.Word:
		data := requestFurtherBytes(2)
		addr = int(binary.LittleEndian.Uint16(data))
	default:
		return nil, fmt.Errorf("unexpected w=%s for mov acc/mem dispatch", w)
	}

	var source interface{}
	var dest interface{}

	switch opcode {
	case fields.MovAccToMem:
		source = acc
		dest = types.MemoryAddress{Address: addr}
	case fields.MovMemToAcc:
		source = types.MemoryAddress{Address: addr}
		dest = acc
	default:
		return nil, fmt.Errorf("unexpected opcode %s for mov acc/mem dispatch", opcode)
	}

	return MovAccMem{
		Opcode: opcode,
		W:      w,
		Source: source,
		Dest:   dest,
	}, nil
}

type MovRegMemSR struct {
	Opcode fields.Opcode
	Mod    fields.Mod
	SR     fields.SR
	Source interface{}
	Dest   interface{}
}

func (mov MovRegMemSR) AsStringInstruction() string {
	return fmt.Sprintf("mov %s, %s", mov.Dest, mov.Source)
}

func DecodeMovRegMemSR(byte1 byte, opcode fields.Opcode, requestFurtherBytes func(int) []byte) (util.InstructionType, error) {
	// | ______  _  _ | __  0 __ ___  | disp-lo | disp-hi |
	// | 100011 0/1 0 | mod 0 SR r/m  | disp-lo | disp-hi |
	byte2 := requestFurtherBytes(1)[0]
	mod, err := fields.DecodeMod(byte2 >> 6)
	if err != nil {
		return nil, err
	}
	sr, err := fields.DecodeSR(byte2 >> 3 & 0x3)
	if err != nil {
		return nil, err
	}

	// always 16 bit address here
	eac, err := EAC(byte2&0x7, mod, fields.Word, requestFurtherBytes)
	if err != nil {
		return nil, fmt.Errorf("cannot calculate EAC for %s: %w", opcode, err)
	}
	fmt.Printf("eac=%s\n", eac.EffectiveAddress)

	var source interface{}
	var dest interface{}

	switch opcode {
	case fields.MovRegMemToSR:
		source = eac.EffectiveAddress
		dest = sr
	case fields.MovSRToRegMem:
		source = sr
		dest = eac.EffectiveAddress
	default:
		return nil, fmt.Errorf("unexpected opcode %s for mov acc/mem dispatch", opcode)
	}

	return MovRegMemSR{
		Opcode: opcode,
		Mod:    mod,
		SR:     sr,
		Source: source,
		Dest:   dest,
	}, nil
}
