package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/temnok/esp32c6/check"
	"github.com/temnok/esp32c6/tap"
	"testing"
)

// "6.1.4. DTM Control and Status (dtmcs, at 0x10)" from "The RISC-V Debug Specification"
func TestTapDtmcs(t *testing.T) {
	defer check.RecoverAndPrintStack(t.Fatal)

	tap.Session(func(conn *tap.Conn) {
		idle := 1 << 12   // 1: Enter Run-Test/Idle and leave it immediately.
		abits := 7 << 4   // 7: The size of address in dmi
		version := 1 << 0 // 1: Version described in spec versions 0.13 and 1.0

		conn.WriteIR(0x10)
		assert.Equal(t, idle|abits|version, conn.ReadDR(32))
	})
}
