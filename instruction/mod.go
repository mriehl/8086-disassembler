package instruction

import (
	"8086-disassembler/util"
	"fmt"
)

type Mod int

const (
	MemoryModeNoDisplacement Mod = iota + 1
	MemoryModeDisplacement8
	MemoryModeDisplacement16
	RegisterToRegister
)

func (mod Mod) String() string {
	switch mod {
	case RegisterToRegister:
		return "RegisterToRegister"
	case MemoryModeNoDisplacement:
		return "MemoryModeNoDisplacement"
	case MemoryModeDisplacement8:
		return "MemoryModeDisplacement8"
	case MemoryModeDisplacement16:
		return "MemoryModeDisplacement16"
	default:
		return "Unknown"
	}
}

func DecodeMod(b byte) (Mod, error) {
	section := b & 0xC0
	mods := map[byte]Mod{
		// 00
		0x00: MemoryModeNoDisplacement,
		// 01
		0x40: MemoryModeDisplacement8,
		// 10
		0x80: MemoryModeDisplacement16,
		// 11
		0xC0: RegisterToRegister,
	}

	mod, ok := mods[section]
	if !ok {
		return 0, fmt.Errorf("unknown mod 0x%X (%s).", b, util.RenderBytes([]byte{b}))
	}
	return mod, nil
}
