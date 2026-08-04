package tests

import (
	"github.com/google/gousb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJtagBypass(t *testing.T) {
	defer handlePanic(t)

	withJtagUsbEndpoints(func(w *gousb.OutEndpoint, r *gousb.InEndpoint) {
		// both registers 0 and 15 are BYPASS
		assert.Equal(t, 0b_01110010, jtagBasicTransaction(w, r, 0x0, 8, 0b_10111001))
		assert.Equal(t, 0b_111111110, jtagBasicTransaction(w, r, 0x1F, 16, 0b_11111111))
	})
}
