package instruction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReg(t *testing.T) {
	// 000
	//   ___
	var lower byte
	lower = 0x0
	wordReg, _ := DecodeReg(lower, Word)
	assert.Equal(t, AX, wordReg)

	byteReg, _ := DecodeReg(lower, Byte)
	assert.Equal(t, AL, byteReg)

	// 011
	//   ___
	lower = 0x3
	wordReg, _ = DecodeReg(lower, Word)
	assert.Equal(t, BX, wordReg)

	byteReg, _ = DecodeReg(lower, Byte)
	assert.Equal(t, BL, byteReg)

	// 100
	//   ___
	lower = 0x4
	wordReg, _ = DecodeReg(lower, Word)
	assert.Equal(t, SP, wordReg)

	byteReg, _ = DecodeReg(lower, Byte)
	assert.Equal(t, AH, byteReg)

	// 111
	//   ___
	lower = 0x7
	wordReg, _ = DecodeReg(lower, Word)
	assert.Equal(t, DI, wordReg)

	byteReg, _ = DecodeReg(lower, Byte)
	assert.Equal(t, BH, byteReg)
}
