package fields

import (
	"8086-disassembler/util"
	"encoding/binary"
	"fmt"
	"strings"
)

type MemoryAddress struct {
	Reg1         Reg
	Reg2         Reg
	Displacement int
}

func (ma MemoryAddress) String() string {
	var buf strings.Builder
	buf.WriteRune('[')
	if ma.Reg1 > 0 {
		buf.WriteString(ma.Reg1.String())
	}
	if ma.Reg2 > 0 {
		buf.WriteString(fmt.Sprintf(" + %s", ma.Reg2.String()))
	}
	if ma.Displacement != 0 {
		if ma.Reg1 == 0 {
			buf.WriteString(fmt.Sprintf("%d", ma.Displacement))
		} else {
			buf.WriteString(fmt.Sprintf(" + %d", ma.Displacement))
		}
	}
	buf.WriteRune(']')

	return buf.String()
}

func DecodeTrailingMemoryLength(rm byte, mod Mod) uint8 {
	switch mod {
	case MemoryModeDisplacement16:
		return 2
	case MemoryModeDisplacement8:
		return 1
	case MemoryModeNoDisplacement:
		// rm=110, mod=00 fuckery, direct word
		if rm == 0x6 {
			return 2
		}
		return 0
	default:
		panic(fmt.Errorf("triny to decode trailing memory for mod %s", mod))
	}
}

func eac(rm byte) (Reg, Reg) {
	switch rm {
	case 0x0:
		return BX, SI
	case 0x1:
		return BX, DI
	case 0x2:
		return BP, SI
	case 0x3:
		return BP, DI
	case 0x4:
		return SI, 0
	case 0x5:
		return DI, 0
	case 0x6:
		return BP, 0
	case 0x7:
		return BX, 0
	default:
		panic(fmt.Errorf("Cannot do EAC for rm=%s", util.RenderByte(rm)))
	}
}

func DecodeMemoryAddress(rm byte, mod Mod, additional []byte) (*MemoryAddress, error) {
	address := MemoryAddress{}

	reg1, reg2 := eac(rm)
	address.Reg1 = reg1
	address.Reg2 = reg2

	switch mod {
	case MemoryModeDisplacement16:
		address.Displacement = int(int16(binary.LittleEndian.Uint16(additional)))
	case MemoryModeDisplacement8:
		address.Displacement = int(int8(additional[0]))
	case MemoryModeNoDisplacement:
		// rm=110, mod=00 fuckery, direct word
		if rm == 0x6 {
			address.Displacement = int(int16(binary.LittleEndian.Uint16(additional)))
			address.Reg1 = 0
			address.Reg2 = 0
		} else {
		}
	default:
		return nil, fmt.Errorf("cannot decode memory address from mod %s", mod)
	}

	return &address, nil

}
