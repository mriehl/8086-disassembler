package instructions

import (
	"8086-disassembler/decoder/fields"
	"8086-disassembler/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeMovRegToFromReg(t *testing.T) {
	raw := []byte{
		// | opcode d w | mod reg r/m |
		// | 100010 0 1 | 11  110 011 |
		util.FromBitstring("10001001"),
		util.FromBitstring("11110011"),
	}

	inst, _ := DecodeMovRmToFromReg(raw[0], fields.MovRmToFromReg, func(i int) []byte {
		return raw[1:]
	})
	assert.Equal(t,
		MovRmToFromRegInstruction{
			Raw:     raw,
			InstBuf: "10001001 11110011",
			Opcode:  fields.MovRmToFromReg,
			Mod:     fields.RegisterMode,
			D:       fields.RegIsSource,
			W:       fields.Word,
			Source:  fields.SI,
			Dest:    fields.BX,
		},
		inst,
	)
}

func TestDecodeMovRegToFromMemoryDispByte(t *testing.T) {
	raw := []byte{
		// | opcode d w | mod reg r/m | disp |
		// | 100010 0 0 | 01  101 110 | 1    |
		util.FromBitstring("10001000"),
		util.FromBitstring("01101110"),
		util.FromBitstring("00000001"),
	}

	firstTime := true

	inst, _ := DecodeMovRmToFromReg(raw[0], fields.MovRmToFromReg, func(i int) []byte {
		if firstTime {
			firstTime = false
			return raw[1:2]
		} else {
			return raw[2:]
		}
	})
	assert.Equal(t,
		MovRmToFromRegInstruction{
			Raw:     raw,
			InstBuf: "10001000 01101110 00000001",
			Opcode:  fields.MovRmToFromReg,
			Mod:     fields.MemoryModeDisplacement8,
			D:       fields.RegIsSource,
			W:       fields.Byte,
			Source:  fields.CH,
			Dest: &fields.MemoryAddress{
				Reg1:         fields.BP,
				Reg2:         0,
				Displacement: 1,
			},
		},
		inst,
	)
}

func TestDecodeMovRegToFromMemoryDispWord(t *testing.T) {
	raw := []byte{
		// | opcode d w | mod reg r/m | displo   | disphi   |
		// | 100010 0 1 | 10  101 110 | 00000001 | 10000000 |
		util.FromBitstring("10001001"),
		util.FromBitstring("10101110"),
		util.FromBitstring("00000001"),
		util.FromBitstring("10000000"),
	}

	firstTime := true

	inst, _ := DecodeMovRmToFromReg(raw[0], fields.MovRmToFromReg, func(i int) []byte {
		if firstTime {
			firstTime = false
			return raw[1:2]
		} else {
			return raw[2:]
		}
	})
	assert.Equal(t,
		MovRmToFromRegInstruction{
			Raw:     raw,
			InstBuf: "10001001 10101110 00000001 10000000",
			Opcode:  fields.MovRmToFromReg,
			Mod:     fields.MemoryModeDisplacement16,
			D:       fields.RegIsSource,
			W:       fields.Word,
			Source:  fields.BP,
			Dest: &fields.MemoryAddress{
				Reg1:         fields.BP,
				Reg2:         0,
				Displacement: 32769,
			},
		},
		inst,
	)
}

func TestDecodeMovImmediateByteToReg(t *testing.T) {
	raw := []byte{
		// | opcode d reg | value    |
		// | 1011   0 001 | 00001100 |
		util.FromBitstring("10110001"),
		util.FromBitstring("00001100"),
	}

	inst, _ := DecodeMovImmediateToReg(raw[0], fields.MovRmToFromReg, func(i int) []byte {
		return raw[1:]
	})
	assert.Equal(t,
		MovImmediateToRegInstruction{
			Raw:         raw,
			InstBuf:     "10110001 00001100",
			Opcode:      fields.MovRmToFromReg,
			W:           fields.Byte,
			SourceValue: 12,
			Dest:        fields.CL,
		},
		inst,
	)
}

func TestDecodeMovImmediateWordToReg(t *testing.T) {
	raw := []byte{
		// | opcode d reg | lo       | hi       |
		// | 1011   1 001 | 00000001 | 10000000 |
		util.FromBitstring("10111001"),
		util.FromBitstring("00000001"),
		util.FromBitstring("10000000"),
	}

	inst, _ := DecodeMovImmediateToReg(raw[0], fields.MovRmToFromReg, func(i int) []byte {
		return raw[1:]
	})
	assert.Equal(t,
		MovImmediateToRegInstruction{
			Raw:         raw,
			InstBuf:     "10111001 00000001 10000000",
			Opcode:      fields.MovRmToFromReg,
			W:           fields.Word,
			SourceValue: 32769,
			Dest:        fields.CX,
		},
		inst,
	)
}
