package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/temnok/esp32c6/jtag"
	"io"
	"testing"
)

func TestJtagIdcode(t *testing.T) {
	defer handlePanic(t)

	jtag.WithUsbConnection(func(usb io.ReadWriter) {
		// https://docs.espressif.com/projects/esp-idf/en/stable/esp32c6/api-guides/jtag-debugging/index.html#codecell3
		const esp32c6Idcode = 0xDC25
		assert.Equal(t, esp32c6Idcode, jtag.BasicTransaction(usb, 0x1, 32, 0))
	})
}
