package instructions

import (
	"8086-disassembler/decoder/fields"
)

type EACResult struct {
	ReadBytes        []byte
	EffectiveAddress interface{}
}

func EAC(rmSection byte, mod fields.Mod, w fields.W, readFurther func(int) []byte) (*EACResult, error) {
	var rm interface{}
	var err error
	readBytes := []byte{}
	if mod == fields.RegisterMode {
		rm, err = fields.DecodeReg(rmSection, w)
		if err != nil {
			return nil, err
		}
	} else {
		additionalRequired := fields.DecodeTrailingMemoryLength(rmSection, mod)
		readBytes = readFurther(int(additionalRequired))
		rm, err = fields.DecodeMemoryAddress(rmSection, mod, readBytes)
		if err != nil {
			return nil, err
		}
	}

	return &EACResult{
		ReadBytes:        readBytes,
		EffectiveAddress: rm,
	}, nil
}
