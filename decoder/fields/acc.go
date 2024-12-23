package fields

import "fmt"

func DecodeAcc(w W) (Reg, error) {
	switch w {
	case Byte:
		return AL, nil
	case Word:
		return AX, nil
	default:
		return 0, fmt.Errorf("unknown accumulator for w=%s", w)
	}

}
