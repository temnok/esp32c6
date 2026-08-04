package tests

import (
	"github.com/google/gousb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJtagIdcode(t *testing.T) {
	defer handlePanic(t)

	withJtagUsbEndpoints(func(w *gousb.OutEndpoint, r *gousb.InEndpoint) {
		// https://docs.espressif.com/projects/esp-idf/en/stable/esp32c6/api-guides/jtag-debugging/index.html#codecell3
		const esp32c6Idcode = 0xDC25
		assert.Equal(t, esp32c6Idcode, jtagBasicTransaction(w, r, 1, 32, 0))
	})
}
