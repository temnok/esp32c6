package esp32c6

import (
	"github.com/temnok/esp32c6/check"
	"go.bug.st/serial"
	"testing"
	"time"
)

func TestReset(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fatal(err)
		}
	}()

	port := check.B(serial.Open("/dev/ttyACM0", &serial.Mode{
		BaudRate: 115200,
	}))

	defer check.Defer(port.Close)

	resetIntoDownloadMode(port)

	time.Sleep(1 * time.Second)

	resetIntoBootingFromFlash(port)
}

// according to "Table 32.4-1. Reset SoC into Download Mode" of "ESP32-C6 Technical Reference Manual"
func resetIntoDownloadMode(port serial.Port) {
	check.A(port.SetDTR(false))
	check.A(port.SetRTS(false))
	check.A(port.SetDTR(true))
	check.A(port.SetRTS(false))
	check.A(port.SetRTS(true))
	check.A(port.SetDTR(false))
	check.A(port.SetRTS(true))
	check.A(port.SetRTS(false))
}

// according to "Table 32.4-2. Reset SoC into Booting from flash" of "ESP32-C6 Technical Reference Manual"
func resetIntoBootingFromFlash(port serial.Port) {
	check.A(port.SetDTR(false))
	check.A(port.SetRTS(false))
	check.A(port.SetRTS(true))
	check.A(port.SetRTS(false))
}
