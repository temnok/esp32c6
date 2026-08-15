package jtag

import (
	"github.com/google/gousb"
	"github.com/temnok/esp32c6/check"
)

func Session(block func(*Conn)) {
	ctx := gousb.NewContext()
	defer check.Call(ctx.Close)

	dev := check.Err1(ctx.OpenDeviceWithVIDPID(0x303A, 0x1001)) // add udev rule if this fails
	defer check.Call(dev.Close)

	cfg := check.Err1(dev.Config(1)) // connect device if this fails
	defer check.Call(cfg.Close)

	intf := check.Err1(cfg.Interface(2, 0))
	defer intf.Close()

	in := check.Err1(intf.InEndpoint(3))
	out := check.Err1(intf.OutEndpoint(2))

	conn := &Conn{
		dev: dev,
		in:  in,
		out: out,
	}

	conn.setDiv(1)

	block(conn)
}
