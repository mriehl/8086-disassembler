package util

import (
	"fmt"
	"strings"
)

func RenderBytes(bytes []byte) string {
	var buffer strings.Builder

	for _, b := range bytes {
		buffer.WriteString(fmt.Sprintf("%08b ", b))
	}
	return buffer.String()
}
