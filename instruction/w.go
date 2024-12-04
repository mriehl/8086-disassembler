package instruction

import (
	"8086-disassembler/util"
	"fmt"
)

type W int

const (
	Word W = iota + 1
	Byte
)

func (w W) String() string {
	switch w {
	case Word:
		return "Word"
	case Byte:
		return "Byte"
	default:
		return "Unknown"
	}
}

func DecodeW(b byte) (W, error) {
	section := b & 0x1
	ws := map[byte]W{
		0x00: Byte,
		0x1:  Word,
	}

	w, ok := ws[section]
	if !ok {
		return 0, fmt.Errorf("unknown W 0x%X (%s).", b, util.RenderBytes([]byte{b}))
	}
	return w, nil
}
