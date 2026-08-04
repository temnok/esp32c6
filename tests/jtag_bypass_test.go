package tests

import (
	"github.com/google/gousb"
	"github.com/stretchr/testify/assert"
	"github.com/temnok/esp32c6/jtag"
	"testing"
)

func TestJtagBypass(t *testing.T) {
	defer handlePanic(t)

	jtag.WithUsbEndpoints(func(r *gousb.InEndpoint, w *gousb.OutEndpoint) {
		// both registers 0 and 15 are BYPASS
		assert.Equal(t, 0b_01110010, jtag.BasicTransaction(w, r, 0x0, 8, 0b_10111001))
		assert.Equal(t, 0b_111111110, jtag.BasicTransaction(w, r, 0x1F, 16, 0b_11111111))
	})
}
