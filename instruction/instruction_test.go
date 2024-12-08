package instruction

import (
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

	instruction, _ := DecodeMovInstruction(func(i int) []byte {
		return raw
	})
	assert.Equal(t,
		MovInstruction{
			Raw:     raw,
			InstBuf: "10001001 11110011",
			Opcode:  Mov,
			Mod:     RegisterToRegister,
			D:       RegIsSource,
			W:       Word,
			Source:  SI,
			Dest:    BX,
		},
		instruction,
	)
}
