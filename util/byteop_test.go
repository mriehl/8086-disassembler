package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderByte(t *testing.T) {
	assert.Equal(t, "00000000", RenderByte(0x0))
	assert.Equal(t, "00000011", RenderByte(0x3))
	assert.Equal(t, "11111111", RenderByte(0xff))
}

func TestRenderBytes(t *testing.T) {
	assert.Equal(t, "00000000 00110001 01100011", RenderBytes([]byte{0x0, 0x31, 0x63}))
}

func TestFromBitstring(t *testing.T) {
	assert.Equal(t, byte(0x0), FromBitstring("00000000"))
	assert.Equal(t, byte(0x3), FromBitstring("00000011"))
	assert.Equal(t, byte(0xff), FromBitstring("11111111"))
}
