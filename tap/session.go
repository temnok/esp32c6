package tap

import (
	"github.com/temnok/esp32c6/jtag"
)

func Session(block func(*Conn)) {
	jtag.Session(func(jtag *jtag.Conn) {
		conn := &Conn{jtag}

		//conn.InitIdle()

		block(conn)
	})
}
