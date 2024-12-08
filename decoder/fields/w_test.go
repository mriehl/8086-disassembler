package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeW(t *testing.T) {
	w, _ := DecodeW(0x0)
	assert.Equal(t, Byte, w)
	w, _ = DecodeW(0x1)
	assert.Equal(t, Word, w)
}

func TestWStringer(t *testing.T) {
	assert.Equal(t, "Word", Word.String())
	assert.Equal(t, "Byte", Byte.String())
}
