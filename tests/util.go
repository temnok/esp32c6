package tests

import (
	"github.com/google/gousb"
	"github.com/temnok/esp32c6/check"
	"runtime/debug"
	"testing"
)

func handlePanic(t *testing.T) {
	if err := recover(); err != nil {
		t.Fatalf("%v\n%s", err, debug.Stack())
	}
}

func withJtagUsbDevice(f func(dev *gousb.Device)) {
	ctx := gousb.NewContext()
	defer check.Call(ctx.Close)

	dev := check.Err1(ctx.OpenDeviceWithVIDPID(0x303A, 0x1001)) // add udev rule if this fails
	defer check.Call(dev.Close)

	f(dev)
}

func withJtagUsbEndpoints(f func(in *gousb.InEndpoint, out *gousb.OutEndpoint)) {
	withJtagUsbDevice(func(dev *gousb.Device) {

		cfg := check.Err1(dev.Config(1)) // connect device if this fails
		defer check.Call(cfg.Close)

		intf := check.Err1(cfg.Interface(2, 0))
		defer intf.Close()

		in := check.Err1(intf.InEndpoint(3))
		out := check.Err1(intf.OutEndpoint(2))

		f(in, out)

	})
}
