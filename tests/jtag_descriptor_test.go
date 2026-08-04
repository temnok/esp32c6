package tests

import (
	"github.com/google/gousb"
	"github.com/stretchr/testify/assert"
	"github.com/temnok/esp32c6/check"
	"github.com/temnok/esp32c6/jtag"
	"testing"
)

func TestJtagDescriptor(t *testing.T) {
	defer handlePanic(t)

	jtag.WithUsbDevice(func(dev *gousb.Device) {

		data := make([]byte, 10)
		n := check.Err1(dev.Control(0x80, 6, 0x2000, 0, data))
		assert.Equal(t, 10, n)

		// according to "Table 32.3-5. JTAG Capability Descriptors" of "ESP32-C6 Technical Reference Manual"
		assert.Equal(t, []byte{
			1,  // JTAG protocol capability structure version
			10, // Total length of JTAG protocol capabilities
			1,  // Type of this struct: 1 for speed capability struct
			8,  // Length of this speed capabilities struct

			4800 % 256, 4800 / 256, // JTAG base clock speed in 10 kHz increments. Note that the maximum TCK speed is half of this value
			1, 0, // Minimum divider value settable by the VEND_JTAG_SETDIV request
			255, 0, // Maximum divider value settable by the VEND_JTAG_SETDIV request
		}, data)

	})
}
