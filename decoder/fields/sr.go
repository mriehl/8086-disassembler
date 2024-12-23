package fields

import (
	"8086-disassembler/util"
	"fmt"
)

type SR int

const (
	ES SR = iota + 1
	CS
	SS
	DS
)

func (sr SR) String() string {
	switch sr {
	case ES:
		return "ES"
	case CS:
		return "CS"
	case SS:
		return "SS"
	case DS:
		return "DS"
	default:
		return "Unknown"
	}
}

func DecodeSR(srSection byte) (SR, error) {
	srs := map[byte]SR{
		0x0: ES,
		0x1: CS,
		0x2: SS,
		0x3: DS,
	}

	w, ok := srs[srSection]
	if !ok {
		return 0, fmt.Errorf("unknown segment register 0x%X (%s).", srSection, util.RenderByte(srSection))
	}
	return w, nil
}
