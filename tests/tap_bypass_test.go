package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/temnok/esp32c6/check"
	"github.com/temnok/esp32c6/tap"
	"testing"
)

func TestTapBypass(t *testing.T) {
	defer check.RecoverAndPrintStack(t.Fatal)

	tap.Session(func(conn *tap.Conn) {
		// both registers 0 and 15 are BYPASS
		conn.WriteIR(0x00)
		assert.Equal(t, 0b_01110010, conn.ScanDR(8, 0b_10111001))

		conn.WriteIR(0x1F)
		assert.Equal(t, 0b_111111110, conn.ScanDR(16, 0b_11111111))
	})
}
