package instruction

import (
	"8086-disassembler/util"
	"fmt"
)

type Reg int

const (
	// W=0
	AL Reg = iota + 1
	CL
	DL
	BL
	AH
	CH
	DH
	BH

	// W=1
	AX
	CX
	DX
	BX
	SP
	BP
	SI
	DI
)

func (reg Reg) String() string {
	switch reg {
	case AL:
		return "AL"
	case CL:
		return " CL"
	case DL:
		return " DL"
	case BL:
		return " BL"
	case AH:
		return " AH"
	case CH:
		return " CH"
	case DH:
		return " DH"
	case BH:
		return " BH"
	case AX:
		return " AX"
	case CX:
		return " CX"
	case DX:
		return " DX"
	case BX:
		return " BX"
	case SP:
		return " SP"
	case BP:
		return " BP"
	case SI:
		return " SI"
	case DI:
		return " DI"
	default:
		return "Unknown"
	}
}

func DecodeReg(regSection byte, w W) (Reg, error) {
	byteRegs := map[byte]Reg{
		0x0: AL,
		0x1: CL,
		0x2: DL,
		0x3: BL,
		0x4: AH,
		0x5: CH,
		0x6: DH,
		0x7: BH,
	}
	wordRegs := map[byte]Reg{
		0x0: AX,
		0x1: CX,
		0x2: DX,
		0x3: BX,
		0x4: SP,
		0x5: BP,
		0x6: SI,
		0x7: DI,
	}

	var reg Reg
	var ok bool
	if w == Byte {
		reg, ok = byteRegs[regSection]
	} else if w == Word {
		reg, ok = wordRegs[regSection]
	} else {
		panic(fmt.Errorf("unknown reg type for W=%v reg=0x%X", w, regSection))
	}

	if !ok {
		return 0, fmt.Errorf("unknown reg 0x%X for W=%v (%s).", regSection, w, util.RenderBytes([]byte{regSection}))
	}
	return reg, nil
}
