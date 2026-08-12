package jtag

import (
	"fmt"
	"github.com/google/gousb"
	"github.com/temnok/esp32c6/check"
)

type Conn struct {
	dev *gousb.Device
	in  *gousb.InEndpoint
	out *gousb.OutEndpoint
}

func (c *Conn) Control(rType, request uint8, val, idx uint16, buf []byte) {
	n := check.Err1(c.dev.Control(rType, request, val, idx, buf))
	if n != len(buf) {
		panic(fmt.Errorf("unexpected JTAG control response: %v, expected %v", n, len(buf)))
	}
}

func (c *Conn) Read(buf []byte) {
	n := check.Err1(c.in.Read(buf))
	if n != len(buf) {
		panic(fmt.Errorf("unexpected JTAG read response: %v, expected %v", n, len(buf)))
	}
}

func (c *Conn) Write(buf []byte) {
	n := check.Err1(c.out.Write(buf))
	if n != len(buf) {
		panic(fmt.Errorf("unexpected JTAG write response: %v, expected %v", n, len(buf)))
	}
}

func (c *Conn) Reset() {
	c.Write([]byte{
		0x22,
		0x22,
		0x20, // 5xTMS==Reset, Idle
	})
}

func (c *Conn) WriteIR(ir int) {
	i := byte(ir)

	c.Write([]byte{
		0x22,                      // Select-DR, Select-IR,
		0x00,                      // Capture-IR, Shift-IR,
		0x00 | i&1<<4 | i>>1&1,    // Shift-IR, Shift-IR,
		0x00 | i>>2&1<<4 | i>>3&1, // Shift-IR, Shift-IR,
		0x22 | i>>4&1<<4,          // Exit1-IR, Update-IR,
		0x00,                      // Idle, Idle
	})
}

func (c *Conn) Transaction(irVal, tdLen, tdi int) (tdo int) {
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

	c.out.Write(nibbles)

	tdoBytes := make([]byte, (tdLen+7)/8)
	c.in.Read(tdoBytes)

	for i, b := range tdoBytes {
		tdo |= int(b) << (i * 8)
	}

	return
}
