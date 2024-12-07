package instruction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeD(t *testing.T) {
	d, _ := DecodeD(0x0)
	assert.Equal(t, d, RegIsSource)

	d, _ = DecodeD(0x1)
	assert.Equal(t, d, RegIsDest)
}

func TestDStringer(t *testing.T) {
	assert.Equal(t, "RegIsSource", RegIsSource.String())
	assert.Equal(t, "RegIsDest", RegIsDest.String())
}
