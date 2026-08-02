package esp32c6

import (
	"github.com/google/gousb"
	"github.com/stretchr/testify/assert"
	"github.com/temnok/esp32c6/check"
	"runtime/debug"
	"testing"
)

func TestGetDescriptor(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fatalf("%v\n%s", err, debug.Stack())
		}
	}()

	ctx := gousb.NewContext()
	defer check.Call(ctx.Close)

	dev := check.E1(ctx.OpenDeviceWithVIDPID(0x303A, 0x1001)) // add udev rule if this fails
	defer check.Call(dev.Close)

	cfg := check.E1(dev.Config(1))
	defer check.Call(cfg.Close)

	intf := check.E1(cfg.Interface(2, 0))
	defer intf.Close()

	outEP := check.E1(intf.OutEndpoint(2))
	inEP := check.E1(intf.InEndpoint(3))

	assert.Equal(t, "ep #2 OUT (address 0x02) bulk [64 bytes]", outEP.String())
	assert.Equal(t, "ep #3 IN (address 0x83) bulk [64 bytes]", inEP.String())
}
