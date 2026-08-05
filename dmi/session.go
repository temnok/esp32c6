package dmi

import (
	"github.com/temnok/esp32c6/jtag"
	"io"
)

type Connection struct {
	usb io.ReadWriter
}

func Session(block func(*Connection)) {
	jtag.WithUsbConnection(func(usb io.ReadWriter) {
		Initialize(usb)

		block(&Connection{usb})
	})
}

func (c *Connection) Read(addr int) int {
	return Read(c.usb, addr)
}

func (c *Connection) Write(addr, data int) {
	Write(c.usb, addr, data)
}
