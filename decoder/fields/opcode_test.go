package fields

import (
	"8086-disassembler/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeOpcode(t *testing.T) {
	opcode, _ := DecodeOpcode(util.FromBitstring("10001000"))
	assert.Equal(t, Mov, opcode)
}
