package fields

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

func DecodeMod(modSection byte) (Mod, error) {
	mods := map[byte]Mod{
		// 00
		0x0: MemoryModeNoDisplacement,
		// 01
		0x1: MemoryModeDisplacement8,
		// 10
		0x2: MemoryModeDisplacement16,
		// 11
		0x3: RegisterToRegister,
	}

	mod, ok := mods[modSection]
	if !ok {
		return 0, fmt.Errorf("unknown mod 0x%X (%s).", modSection, util.RenderByte(modSection))
	}
	return mod, nil
}
