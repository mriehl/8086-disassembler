package instructions

import (
	"8086-disassembler/decoder/fields"
	"8086-disassembler/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeMov(t *testing.T) {
	raw := []byte{
		// | ______ _ _ | __  ___ ___ |
		// | opcode d w | mod reg r/m |
		// | 100010 0 1 | 11  110 011 |
		util.FromBitstring("10001001"),
		util.FromBitstring("11110011"),
	}

	inst, _ := DecodeMovInstruction(fields.Mov, func(i int) []byte {
		return raw
	})
	assert.Equal(t,
		MovInstruction{
			Raw:     raw,
			InstBuf: "10001001 11110011",
			Opcode:  fields.Mov,
			Mod:     fields.RegisterToRegister,
			D:       fields.RegIsSource,
			W:       fields.Word,
			Source:  fields.SI,
			Dest:    fields.BX,
		},
		inst,
	)
}
