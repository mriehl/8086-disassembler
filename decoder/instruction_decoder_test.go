package decoder

import (
	"8086-disassembler/util"
	"bufio"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func assertOneResult(t *testing.T, expected DecodeResult, inst ...byte) {
	input := []byte(inst)
	reader := bufio.NewReader(strings.NewReader(string(input)))

	ch := DecodeInstructions(reader)
	select {
	case result := <-ch:
		assert.Equal(t, expected, result)
	case <-time.After(1 * time.Second):
		assert.Fail(t, "Timeout waiting for result from channel")
	}
	_, ok := <-ch
	assert.False(t, ok, "Channel should be closed after one result")
}

func TestBrokenOpcode(t *testing.T) {
	assertOneResult(t,
		DecodeResult{
			Value: nil,
			Error: errors.New("error while decoding instruction (read=00000000) due to unknown opcode 0x0 (00000000)."),
		},
		util.FromBitstring("00000000"),
	)
}
