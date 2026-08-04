package tests

import (
	"github.com/google/gousb"
	"github.com/stretchr/testify/assert"
	"github.com/temnok/esp32c6/jtag"
	"testing"
)

// "6.1.4. DTM Control and Status (dtmcs, at 0x10)" from "The RISC-V Debug Specification"
func TestJtagDtmcs(t *testing.T) {
	defer handlePanic(t)

	jtag.WithUsbEndpoints(func(w *gousb.OutEndpoint, r *gousb.InEndpoint) {
		idle := 1 << 12   // 1: Enter Run-Test/Idle and leave it immediately.
		abits := 7 << 4   // 7: The size of address in dmi
		version := 1 << 0 // 1: Version described in spec versions 0.13 and 1.0

		assert.Equal(t, idle|abits|version, jtag.BasicTransaction(w, r, 0x10, 32, 0))
	})
}
