package dmi

import (
	"fmt"
	"github.com/temnok/esp32c6/tap"
)

type Conn struct {
	tap *tap.Conn
}

func (conn *Conn) initialize() {
	conn.tap.InitIdle()
	conn.tap.WriteIR(0x11)

	conn.Write(Dmcontrol, 0<<DmcontrolDmactive)
	for conn.Read(Dmcontrol)>>DmcontrolDmactive&1 != 0 {
	}

	conn.Write(Dmcontrol, 1<<DmcontrolDmactive)
	for conn.Read(Dmcontrol)>>DmcontrolDmactive&1 == 0 {
	}
}

func (c *Conn) Read(addr int) int {
	val := c.tap.WriteAndReadDR(41, addr<<34|1, 4)
	a, d, op := val>>34, int(uint32(val>>2)), val&3

	if op != 0 {
		panic(fmt.Errorf("reading DMI returned %v", op))
	}

	if a != addr {
		panic(fmt.Errorf("reading DMI returned address 0x%X, expected 0x%X", a, addr))
	}

	return d
}

func (c *Conn) Write(addr, data int) {
	c.tap.WriteDR(41, addr<<34|data<<2|2)
}
