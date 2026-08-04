package tests

import (
	"github.com/google/gousb"
	"github.com/stretchr/testify/assert"
	"github.com/temnok/esp32c6/jtag"
	"testing"
)

func TestJtagEndpoint(t *testing.T) {
	defer handlePanic(t)

	jtag.WithUsbEndpoints(func(w *gousb.OutEndpoint, r *gousb.InEndpoint) {
		assert.Equal(t, "ep #2 OUT (address 0x02) bulk [64 bytes]", w.String())
		assert.Equal(t, "ep #3 IN (address 0x83) bulk [64 bytes]", r.String())
	})
}
