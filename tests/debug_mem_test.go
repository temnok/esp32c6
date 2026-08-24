package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/temnok/esp32c6/check"
	"github.com/temnok/esp32c6/debug"
	"testing"
)

func TestMemAccess(t *testing.T) {
	defer check.RecoverAndPrintStack(t.Fatal)

	debug.Session(func(conn *debug.Conn) {
		conn.HartResetAndHalt(0)

		data := []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE, 0xF0}
		conn.WriteMem(0x4080_0000, data)

		mem := make([]byte, len(data))
		conn.ReadMem(0x4080_0000, mem)

		assert.Equal(t, data, mem)
	})
}
