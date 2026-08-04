package tests

import (
	"github.com/google/gousb"
	"github.com/stretchr/testify/assert"
	"github.com/temnok/esp32c6/jtag"
	"testing"
)

func TestJtagIdcode(t *testing.T) {
	defer handlePanic(t)

	jtag.WithUsbEndpoints(func(w *gousb.OutEndpoint, r *gousb.InEndpoint) {
		// https://docs.espressif.com/projects/esp-idf/en/stable/esp32c6/api-guides/jtag-debugging/index.html#codecell3
		const esp32c6Idcode = 0xDC25
		assert.Equal(t, esp32c6Idcode, jtag.BasicTransaction(w, r, 0x1, 32, 0))
	})
}
