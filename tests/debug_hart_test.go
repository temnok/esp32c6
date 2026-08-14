package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/temnok/esp32c6/check"
	"github.com/temnok/esp32c6/debug"
	"testing"
)

func TestHartCount(t *testing.T) {
	defer check.RecoverAndPrintStack(t.Fatal)

	debug.Session(func(conn *debug.Conn) {
		assert.Equal(t, 2, conn.HartCount())

		conn.HartHalt(0)
		//conn.HartHalt(1)
	})
}
