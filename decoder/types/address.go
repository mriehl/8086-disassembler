package types

import "fmt"

type MemoryAddress struct {
	Address int
}

func (ma MemoryAddress) String() string {
	return fmt.Sprintf("[%d]", ma.Address)
}
