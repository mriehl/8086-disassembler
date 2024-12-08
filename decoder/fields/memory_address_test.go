package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertResultsIn(t *testing.T, rm byte, mod Mod, additional []byte, expectedInst string) {
	address, err := DecodeMemoryAddress(
		rm,
		mod,
		additional,
	)
	assert.NoError(t, err)
	assert.Equal(t, expectedInst, address.String())
}

func TestDecodeToStringReg1(t *testing.T) {
	assertResultsIn(t,
		0x4,
		MemoryModeNoDisplacement,
		[]byte{},
		"[SI]")
}
func TestDecodeToStringReg1DispByte(t *testing.T) {
	assertResultsIn(t,
		0x4,
		MemoryModeDisplacement8,
		[]byte{0x3},
		"[SI + 3]")
}
func TestDecodeToStringReg1DispWord(t *testing.T) {
	assertResultsIn(t,
		0x5,
		MemoryModeDisplacement16,
		[]byte{0x4, 0x2},
		"[DI + 516]")
}

func TestDecodeToStringReg1Reg2(t *testing.T) {
	assertResultsIn(t,
		0x1,
		MemoryModeNoDisplacement,
		[]byte{},
		"[BX + DI]")
}
func TestDecodeToStringReg1Reg2DispByte(t *testing.T) {
	assertResultsIn(t,
		0x3,
		MemoryModeDisplacement8,
		[]byte{0x6},
		"[BP + DI + 6]")
}
func TestDecodeToStringReg1Reg2DispWord(t *testing.T) {
	assertResultsIn(t,
		0x2,
		MemoryModeDisplacement16,
		[]byte{0x6, 0x1},
		"[BP + SI + 262]")
}

func TestDecodeToStringDirect(t *testing.T) {
	assertResultsIn(t,
		0x6,
		MemoryModeNoDisplacement,
		[]byte{0x6, 0x1},
		"[262]")
}
