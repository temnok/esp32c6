package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/temnok/esp32c6/jtag"
	"io"
	"testing"
)

func TestDmiDmstatus(t *testing.T) {
	defer handlePanic(t)

	jtag.WithUsbConnection(func(usb io.ReadWriter) {
		jtag.SelectDMI(usb)

		addr, dmstatus, op := jtag.AccessDMI(usb, 0x11, 0, jtag.DmiOpRead)
		assert.Equal(t, 0, op)
		assert.Equal(t, 2, dmstatus&0xF) // 2 (0.13): There is a Debug Module and it conforms to version 0.13 of this specification
		assert.Equal(t, 0x11, addr)
	})
}
