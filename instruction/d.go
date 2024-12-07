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

func DecodeD(dSection byte) (D, error) {
	ds := map[byte]D{
		0x0: RegIsSource,
		0x1: RegIsDest,
	}

	d, ok := ds[dSection]
	if !ok {
		return 0, fmt.Errorf("unknown D 0x%X (%s).", dSection, util.RenderBytes([]byte{dSection}))
	}
	return d, nil
}
