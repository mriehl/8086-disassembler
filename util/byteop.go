package util

import (
	"fmt"
	"strconv"
	"strings"
)

func RenderByte(byte byte) string {
	return fmt.Sprintf("%08b", byte)
}

func RenderBytes(bytes []byte) string {
	var buffer strings.Builder

	for i, b := range bytes {
		buffer.WriteString(RenderByte(b))
		if i != len(bytes)-1 {
			buffer.WriteRune(' ')
		}
	}
	return buffer.String()
}

func FromBitstring(s string) byte {
	var str string
	l := len(s)
	if l > 8 {
		panic(fmt.Errorf("bitstring too long: '%s'", s))
	}

	if l-8 < 0 {
		str = string(s[0:l])
	} else {
		str = string(s[l-8 : l])
	}
	v, err := strconv.ParseUint(str, 2, 8)
	if err != nil {
		panic(err)
	}
	return byte(v)
}
