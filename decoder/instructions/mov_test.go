package instructions

import (
	"8086-disassembler/decoder/fields"
	"8086-disassembler/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeMovRmToFromReg(t *testing.T) {
	raw := []byte{
		// | opcode d w | mod reg r/m |
		// | 100010 0 1 | 11  110 011 |
		util.FromBitstring("10001001"),
		util.FromBitstring("11110011"),
	}

	inst, _ := DecodeMovRmToFromReg(raw[0], fields.MovRmToFromReg, func(i int) []byte {
		return raw
	})
	assert.Equal(t,
		MovRmToFromRegInstruction{
			Raw:     raw,
			InstBuf: "10001001 11110011",
			Opcode:  fields.MovRmToFromReg,
			Mod:     fields.RegisterToRegister,
			D:       fields.RegIsSource,
			W:       fields.Word,
			Source:  fields.SI,
			Dest:    fields.BX,
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
		return raw
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
		return raw
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
