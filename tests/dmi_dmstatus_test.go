package tests

import (
	"github.com/temnok/esp32c6/jtag"
	"io"
	"testing"
)

func xTestDmiDmstatus(t *testing.T) {
	defer handlePanic(t)

	jtag.WithUsbConnection(func(usb io.ReadWriter) {
		//data, op := jtag.AccessDMI(usb, 0x11<<34|0<<2|1)
		//assert.Equal(t, 0, op)
		//fmt.Printf("dmstatus = %032b\naddr = %07b\n", dmstatus, addr)
	})
}
