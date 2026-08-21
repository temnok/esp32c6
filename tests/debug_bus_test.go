package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/temnok/esp32c6/check"
	"github.com/temnok/esp32c6/debug"
	"testing"
)

func TestBusAccess(t *testing.T) {
	defer check.RecoverAndPrintStack(t.Fatal)

	debug.Session(func(conn *debug.Conn) {
		conn.HartResetAndHalt(0)

		data := []uint32{0x12345678, 0x11223344, 0x55667788, 0x87654321}
		conn.WriteBus(0x4080_0000, data)

		mem := make([]uint32, 4)
		conn.ReadBus(0x4080_0000, mem)

		assert.Equal(t, data, mem)
	})
}
