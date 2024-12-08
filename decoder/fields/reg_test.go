package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeReg(t *testing.T) {
	var lower byte
	lower = 0x0
	wordReg, _ := DecodeReg(lower, Word)
	assert.Equal(t, AX, wordReg)

	byteReg, _ := DecodeReg(lower, Byte)
	assert.Equal(t, AL, byteReg)

	lower = 0x3
	wordReg, _ = DecodeReg(lower, Word)
	assert.Equal(t, BX, wordReg)

	byteReg, _ = DecodeReg(lower, Byte)
	assert.Equal(t, BL, byteReg)

	lower = 0x4
	wordReg, _ = DecodeReg(lower, Word)
	assert.Equal(t, SP, wordReg)

	byteReg, _ = DecodeReg(lower, Byte)
	assert.Equal(t, AH, byteReg)

	lower = 0x7
	wordReg, _ = DecodeReg(lower, Word)
	assert.Equal(t, DI, wordReg)

	byteReg, _ = DecodeReg(lower, Byte)
	assert.Equal(t, BH, byteReg)
}

func TestRegStringer(t *testing.T) {
	assert.Equal(t, CX.String(), "CX")

}
