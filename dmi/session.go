package dmi

import (
	"github.com/temnok/esp32c6/tap"
)

func Session(block func(*Conn)) {
	tap.Session(func(tap *tap.Conn) {
		conn := &Conn{tap}

		conn.initialize()

		block(conn)
	})
}
