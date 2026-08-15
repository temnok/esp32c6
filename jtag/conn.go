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

	nibbles       []byte
	pendingNibble bool
}

func (c *Conn) setDiv(div int) {
	c.Control(0x40, 0, div, 0, nil) // VEND_JTAG_SETDIV
}

func (c *Conn) Control(requestType, request, value, index int, buf []byte) {
	n := check.Err1(c.dev.Control(byte(requestType), byte(request), uint16(value), uint16(index), buf))
	if n != len(buf) {
		panic(fmt.Errorf("unexpected JTAG control response: %v, expected %v", n, len(buf)))
	}
}

func (c *Conn) ClockCapTmsTdi(cap, tms, tdi int) {
	c.appendNibble(cap&1<<2 | tms&1<<1 | tdi&1) // CMD_CLK
}

func (c *Conn) ClockTms(tms int) {
	c.ClockCapTmsTdi(0, tms, 0)
}

func (c *Conn) Repeat(n int) {
	for i := n; i > 0; i >>= 2 {
		c.appendNibble(0b_1100 | i&3) // CMD_REP
	}
}

func (c *Conn) Flush(out []byte) {
	if len(out) > 0 {
		c.appendNibble(0b_1010) // CMD_FLUSH
	}

	if len(c.nibbles) > 0 {
		c.write(c.nibbles)
		c.nibbles = nil
		c.pendingNibble = false
	}

	if len(out) > 0 {
		c.read(out)
	}
}

func (c *Conn) appendNibble(nibble int) {
	if c.pendingNibble {
		c.nibbles[len(c.nibbles)-1] |= byte(nibble)
	} else {
		c.nibbles = append(c.nibbles, byte(nibble)<<4)
	}

	c.pendingNibble = !c.pendingNibble
}

func (c *Conn) read(buf []byte) {
	n := check.Err1(c.in.Read(buf))
	if n != len(buf) {
		panic(fmt.Errorf("unexpected JTAG read response: %v, expected %v", n, len(buf)))
	}
}

func (c *Conn) write(buf []byte) {
	n := check.Err1(c.out.Write(buf))
	if n != len(buf) {
		panic(fmt.Errorf("unexpected JTAG write response: %v, expected %v", n, len(buf)))
	}
}
