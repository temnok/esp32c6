package jtag

import (
	"fmt"
	"github.com/temnok/esp32c6/check"
	"io"
)

func SelectDMI(usb io.Writer) {
	// write 0x11 to IR
	nibbles := []byte{
		0b_0010_0010,
		0b_0010_0010,
		0b_0010_0000, // Reset, Idle,
		0b_0010_0010, // Select-DR, Select-IR,
		0b_0000_0000, // Capture-IR, Shift-IR,
		0b_0001_0000, // Shift-IR(1), Shift-IR(0),
		0b_0000_0000, // Shift-IR(0), Shift-IR(0),
		0b_0011_0010, // Exit1-IR(1), Update-IR,
		0b_0000_0000, // Idle, Idle
	}

	n := check.Err1(usb.Write(nibbles))
	if n != len(nibbles) {
		panic(fmt.Errorf("SelectDMI/usb.Write() returned %v, expected %v", n, len(nibbles)))
	}
}

// Debug Module Interface Access (dmi, at 0x11)
func AccessDMI(usb io.ReadWriter, addr, data, op int) (int, int) {
	nibbles := []byte{
		0x20,                                   // Select-DR, Capture-DR,
		0x00 | byte(op&1),                      // Shift-DR, Shift-DR,
		0x00 | byte(op>>1&1)<<4 | byte(data&1), // Shift-DR, Shift-DR,
		0x00 | byte(data>>1&1)<<4 | byte(data>>2&1),   // Shift-DR, Shift-DR,
		0x00 | byte(data>>3&1)<<4 | byte(data>>4&1),   // Shift-DR, Shift-DR,
		0x00 | byte(data>>5&1)<<4 | byte(data>>6&1),   // Shift-DR, Shift-DR,
		0x00 | byte(data>>7&1)<<4 | byte(data>>8&1),   // Shift-DR, Shift-DR,
		0x00 | byte(data>>9&1)<<4 | byte(data>>10&1),  // Shift-DR, Shift-DR,
		0x00 | byte(data>>11&1)<<4 | byte(data>>12&1), // Shift-DR, Shift-DR,
		0x00 | byte(data>>13&1)<<4 | byte(data>>14&1), // Shift-DR, Shift-DR,
		0x00 | byte(data>>15&1)<<4 | byte(data>>16&1), // Shift-DR, Shift-DR,
		0x00 | byte(data>>17&1)<<4 | byte(data>>18&1), // Shift-DR, Shift-DR,
		0x00 | byte(data>>19&1)<<4 | byte(data>>20&1), // Shift-DR, Shift-DR,
		0x00 | byte(data>>21&1)<<4 | byte(data>>22&1), // Shift-DR, Shift-DR,
		0x00 | byte(data>>23&1)<<4 | byte(data>>24&1), // Shift-DR, Shift-DR,
		0x00 | byte(data>>25&1)<<4 | byte(data>>26&1), // Shift-DR, Shift-DR,
		0x00 | byte(data>>27&1)<<4 | byte(data>>28&1), // Shift-DR, Shift-DR,
		0x00 | byte(data>>29&1)<<4 | byte(data>>30&1), // Shift-DR, Shift-DR,
		0x00 | byte(data>>31&1)<<4 | byte(addr&1),     // Shift-DR, Shift-DR,
		0x00 | byte(addr>>1&1)<<4 | byte(addr>>2&1),   // Shift-DR, Shift-DR,
		0x00 | byte(addr>>3&1)<<4 | byte(addr>>4&1),   // Shift-DR, Shift-DR,
		0x02 | byte(addr>>5&1)<<4 | byte(addr>>6&1),   // Shift-DR, Exit1-DR,
		0x20, // Update-DR, Idle,
		0x20, // Select-DR, Capture-DR,
		0x04, // Shift-DR, Shift-DR with capture,
		0x44, // Shift-DR with capture, Shift-DR with capture,

	}

	_ = nibbles

	return 0, 0
}
