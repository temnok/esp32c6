package jtag

import (
	"io"
)

// Debug Module Interface Access (dmi, at 0x11)
func DMIAccess(usb io.ReadWriter, input int) int {
	// 41 = 7 bit address + 32 bit data + 2 bit op
	return BasicTransaction(usb, 0x11, 41, input)
}
