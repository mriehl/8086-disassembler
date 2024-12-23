package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByte(t *testing.T) {
	reg, err := DecodeAcc(Byte)
	assert.NoError(t, err)
	assert.Equal(t, AL, reg)
}

func TestWord(t *testing.T) {
	reg, err := DecodeAcc(Word)
	assert.NoError(t, err)
	assert.Equal(t, AX, reg)
}
