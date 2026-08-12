package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/temnok/esp32c6/check"
	"github.com/temnok/esp32c6/jtag"
	"testing"
)

func TestJtagBypass(t *testing.T) {
	defer check.RecoverAndPrintStack(t.Fatal)

	jtag.Session(func(conn *jtag.Conn) {
		// both registers 0 and 15 are BYPASS
		assert.Equal(t, 0b_01110010, conn.Transaction(0x0, 8, 0b_10111001))
		assert.Equal(t, 0b_111111110, conn.Transaction(0x1F, 16, 0b_11111111))
	})
}
