package tap

import (
	"github.com/temnok/esp32c6/jtag"
)

type Conn struct {
	jtag *jtag.Conn
}

func (c *Conn) InitIdle() {
	c.jtag.ClockTms(1)
	c.jtag.Repeat(4)
	c.jtag.ClockTms(0) // Idle
}

func (c *Conn) WriteIR(val int) {
	c.jtag.ClockTms(1) // Select-DR
	c.scan(0, 5, val)
	c.jtag.Flush(nil)
}

func (c *Conn) ReadDR(bitLen int) int {
	c.jtag.ClockTms(1)             // Select-DR
	c.jtag.ClockTms(0)             // Capture-DR
	c.jtag.ClockTms(0)             // Shift-DR
	c.jtag.ClockCapTmsTdi(1, 0, 0) // Shift-DR

	if bitLen > 2 {
		c.jtag.Repeat(bitLen - 2)
	}

	c.jtag.ClockCapTmsTdi(1, 1, 0) // Exit1-DR
	c.jtag.ClockTms(1)             // Update-DR
	c.jtag.ClockTms(0)             // Idle

	buf := make([]byte, (bitLen+7)/8)
	c.jtag.Flush(buf)

	out := 0
	for i, b := range buf {
		out |= int(b) << (i * 8)
	}

	return out
}

func (c *Conn) WriteDR(len, val int) {
	c.scan(0, len, val)
	c.jtag.Flush(nil)
}

func (c *Conn) ScanDR(len, val int) int {
	c.scan(1, len, val)

	buf := make([]byte, (len+7)/8)
	c.jtag.Flush(buf)

	out := 0
	for i, b := range buf {
		out |= int(b) << (i * 8)
	}

	return out
}

func (c *Conn) WriteAndReadDR(len, val, idleCount int) int {
	c.scan(0, len, val)

	if idleCount > 0 {
		c.jtag.Repeat(idleCount)
	}

	return c.ReadDR(len)
}

func (c *Conn) scan(cap, len, val int) {
	c.jtag.ClockTms(1) // Select-xR
	c.jtag.ClockTms(0) // Capture-xR
	c.jtag.ClockTms(0) // Shift-xR

	for i := range len - 1 {
		c.jtag.ClockCapTmsTdi(cap, 0, val>>i) // Shift-xR
	}

	c.jtag.ClockCapTmsTdi(cap, 1, val>>(len-1)) // Exit1-xR
	c.jtag.ClockTms(1)                          // Update-xR
	c.jtag.ClockTms(0)                          // Idle
}
