package esp32c6

import (
	"github.com/google/gousb"
	"github.com/stretchr/testify/assert"
	"github.com/temnok/esp32c6/check"
	"runtime/debug"
	"testing"
)

func TestJtagDescriptor(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fatalf("%v\n%s", err, debug.Stack())
		}
	}()

	ctx := gousb.NewContext()
	defer check.Call(ctx.Close)

	dev := check.E1(ctx.OpenDeviceWithVIDPID(0x303A, 0x1001)) // add udev rule if this fails
	defer check.Call(dev.Close)

	data := make([]byte, 10)
	n := check.E1(dev.Control(0x80, 6, 0x2000, 0, data))
	assert.Equal(t, 10, n)

	// according to "Table 32.3-5. JTAG Capability Descriptors" of "ESP32-C6 Technical Reference Manual"
	assert.Equal(t, []byte{
		1,                      // JTAG protocol capability structure version
		10,                     // Total length of JTAG protocol capabilities
		1,                      // Type of this struct: 1 for speed capability struct
		8,                      // Length of this speed capabilities struct
		4800 % 256, 4800 / 256, // JTAG base clock speed in 10 kHz increments. Note that the maximum TCK speed is half of this value
		1, 0, // Minimum divider value settable by the VEND_JTAG_SETDIV request
		255, 0, // Maximum divider value settable by the VEND_JTAG_SETDIV request
	}, data)
}
