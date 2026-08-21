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

		conn.WriteBus(0x4080_0000, []int32{1, 2, 3, 4})

		data := make([]int32, 4)
		conn.ReadBus(0x4080_0000, data)

		assert.Equal(t, []int32{1, 2, 3, 4}, data)
	})
}
