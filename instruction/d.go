package instruction

import (
	"8086-disassembler/util"
	"fmt"
)

type D int

const (
	RegIsSource = iota + 1
	RegIsDest
)

func (d D) String() string {
	switch d {
	case RegIsSource:
		return "RegIsSource"
	case RegIsDest:
		return "RegIsDest"
	default:
		return "Unknown"
	}
}

func DecodeD(b byte) (D, error) {
	section := b & 0x2
	ds := map[byte]D{
		0x0: RegIsSource,
		0x2: RegIsDest,
	}

	d, ok := ds[section]
	if !ok {
		return 0, fmt.Errorf("unknown D 0x%X (%s).", b, util.RenderBytes([]byte{b}))
	}
	return d, nil
}
