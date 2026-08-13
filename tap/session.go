package tap

import (
	"github.com/temnok/esp32c6/jtag"
)

func Session(block func(*Conn)) {
	jtag.Session(func(conn *jtag.Conn) {
		block(&Conn{
			conn,
		})
	})
}
