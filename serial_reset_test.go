package esp32c6

import (
	"github.com/temnok/esp32c6/check"
	"go.bug.st/serial"
	"runtime/debug"
	"testing"
	"time"
)

func TestReset(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fatalf("%v\n%s", err, debug.Stack())
		}
	}()

	port := check.E1(serial.Open("/dev/ttyACM0", &serial.Mode{}))

	defer check.Call(port.Close)

	resetIntoDownloadMode(port)

	time.Sleep(100 * time.Millisecond)

	resetIntoBootingFromFlash(port)
}

// according to "Table 32.4-1. Reset SoC into Download Mode" of "ESP32-C6 Technical Reference Manual"
func resetIntoDownloadMode(port serial.Port) {
	check.E(port.SetDTR(false))
	check.E(port.SetRTS(false))
	check.E(port.SetDTR(true))
	check.E(port.SetRTS(false))
	check.E(port.SetRTS(true))
	check.E(port.SetDTR(false))
	check.E(port.SetRTS(true))
	check.E(port.SetRTS(false))
}

// according to "Table 32.4-2. Reset SoC into Booting from flash" of "ESP32-C6 Technical Reference Manual"
func resetIntoBootingFromFlash(port serial.Port) {
	check.E(port.SetDTR(false))
	check.E(port.SetRTS(false))
	check.E(port.SetRTS(true))
	check.E(port.SetRTS(false))
}
