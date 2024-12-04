package instruction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReg(t *testing.T) {
	// 00000000
	//   ___
	var lower byte
	lower = 0x0
	wordReg, _ := DecodeReg(lower, Word)
	assert.Equal(t, AX, wordReg)

	byteReg, _ := DecodeReg(lower, Byte)
	assert.Equal(t, AL, byteReg)

	// 11011001
	//   ___
	lower = 0xD9
	wordReg, _ = DecodeReg(lower, Word)
	assert.Equal(t, BX, wordReg)

	byteReg, _ = DecodeReg(lower, Byte)
	assert.Equal(t, BL, byteReg)

	// 11111001
	//   ___
	lower = 0xFF
	wordReg, _ = DecodeReg(lower, Word)
	assert.Equal(t, DI, wordReg)

	byteReg, _ = DecodeReg(lower, Byte)
	assert.Equal(t, BH, byteReg)
}
