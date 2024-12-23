package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeSR(t *testing.T) {
	sr, _ := DecodeSR(0x0)
	assert.Equal(t, ES, sr)
	sr, _ = DecodeSR(0x1)
	assert.Equal(t, CS, sr)
	sr, _ = DecodeSR(0x2)
	assert.Equal(t, SS, sr)
	sr, _ = DecodeSR(0x3)
	assert.Equal(t, DS, sr)
}

func TestSRStringer(t *testing.T) {
	assert.Equal(t, "ES", ES.String())
	assert.Equal(t, "CS", CS.String())
	assert.Equal(t, "SS", SS.String())
	assert.Equal(t, "DS", DS.String())
}
