package jtag

import (
	"github.com/google/gousb"
	"github.com/temnok/esp32c6/check"
)

func WithUsbDevice(f func(*gousb.Device)) {
	ctx := gousb.NewContext()
	defer check.Call(ctx.Close)

	dev := check.Err1(ctx.OpenDeviceWithVIDPID(0x303A, 0x1001)) // add udev rule if this fails
	defer check.Call(dev.Close)

	f(dev)
}

func WithUsbEndpoints(f func(*gousb.OutEndpoint, *gousb.InEndpoint)) {
	WithUsbDevice(func(dev *gousb.Device) {

		cfg := check.Err1(dev.Config(1)) // connect device if this fails
		defer check.Call(cfg.Close)

		intf := check.Err1(cfg.Interface(2, 0))
		defer intf.Close()

		w := check.Err1(intf.OutEndpoint(2))
		r := check.Err1(intf.InEndpoint(3))

		f(w, r)

	})
}
