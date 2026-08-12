package dmi

import (
	"github.com/temnok/esp32c6/jtag"
)

type Conn struct {
	conn *jtag.Conn
}

func (c *Conn) Read(addr int) int {
	return read(c.conn, addr)
}

func (c *Conn) Write(addr, data int) {
	write(c.conn, addr, data)
}
