package jtag

import (
	"fmt"
	"github.com/temnok/esp32c6/check"
	"io"
)

func BasicTransaction(usb io.ReadWriter, irVal, tdLen, tdi int) (tdo int) {
	ir := byte(irVal) // 5 bits used only

	// according to "Table 32.3-3. Commands of a Nibble" of "ESP32-C6 Technical Reference Manual"
	// CMD_CLK   0 cap tms tdi
	// CMD_FLUSH 1   0   1   0
	// so, 1 = TDI bit set, 2 = TMS bit set, 4 = capture TDO
	nibbles := []byte{
		0, 2, 2, 0, // Idle, Select-DR, Select-IR, Capture-IR
		0, ir & 1, ir >> 1 & 1, ir >> 2 & 1, ir >> 3 & 1, 2 | ir>>4&1, 2, // 5 x Shift-IR, Exit1-IR, Update-IR
		0, 2, 0, 0, // Idle, Select-DR, Capture-DR, Shift-DR
	}

	for i := 0; i < tdLen; i++ {
		nibbles = append(nibbles, byte(4|tdi>>i&1)) // Shift-DR with capture
	}

	nibbles[len(nibbles)-1] |= 2            // Exit1-DR
	nibbles = append(nibbles, 2, 0, 10, 10) // Update-DR, Idle, 2 x Flush

	for i := 0; i*2+1 < len(nibbles); i++ {
		nibbles[i] = nibbles[i*2]<<4 | nibbles[i*2+1] // compress nibbles to two nibbles in a byte
	}

	nibbles = nibbles[:len(nibbles)/2]

	n := check.Err1(usb.Write(nibbles))
	if n != len(nibbles) {
		panic(fmt.Errorf("usb.Write() returned %v, expected %v", n, len(nibbles)))
	}

	tdoBytes := make([]byte, (tdLen+7)/8)
	n = check.Err1(usb.Read(tdoBytes))
	if n != len(tdoBytes) {
		panic(fmt.Errorf("usb.Read() returned %v, expected %v", n, len(tdoBytes)))
	}

	for i, b := range tdoBytes {
		tdo |= int(b) << (i * 8)
	}

	return
}
