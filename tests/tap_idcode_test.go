package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/temnok/esp32c6/check"
	"github.com/temnok/esp32c6/tap"
	"testing"
)

func TestJtagIdcode(t *testing.T) {
	defer check.RecoverAndPrintStack(t.Fatal)

	tap.Session(func(conn *tap.Conn) {
		// https://docs.espressif.com/projects/esp-idf/en/stable/esp32c6/api-guides/jtag-debugging/index.html#codecell3
		conn.WriteIR(0x01)
		assert.Equal(t, 0xDC25, conn.ReadDR(32))
	})
}
