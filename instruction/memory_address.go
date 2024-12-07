package instruction

import "fmt"

type MemoryAddress int

const (
	// MOD != 11
	TODO MemoryAddress = iota + 1
)

func (memoryAddress MemoryAddress) String() string {
	switch memoryAddress {
	case TODO:
		return "MEMORY_TODO"
	default:
		return "Unknown"
	}
}

func DecodeMemoryAddress(b byte, mod Mod) (MemoryAddress, error) {
	if mod == RegisterToRegister {
		panic(fmt.Errorf("cannot decode memory address from mod %s", mod))
	} else {
		return TODO, nil

	}
}
