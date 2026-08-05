package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/temnok/esp32c6/jtag"
	"io"
	"testing"
)

func TestJtagBypass(t *testing.T) {
	defer handlePanic(t)

	jtag.WithUsbConnection(func(usb io.ReadWriter) {
		// both registers 0 and 15 are BYPASS
		assert.Equal(t, 0b_01110010, jtag.BasicTransaction(usb, 0x0, 8, 0b_10111001))
		assert.Equal(t, 0b_111111110, jtag.BasicTransaction(usb, 0x1F, 16, 0b_11111111))
	})
}
