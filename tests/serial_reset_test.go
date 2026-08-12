package tests

import (
	"github.com/temnok/esp32c6/check"
	"go.bug.st/serial"
	"testing"
	"time"
)

func TestSerialReset(t *testing.T) {
	defer check.RecoverAndPrintStack(t.Fatal)

	port := check.Err1(serial.Open("/dev/ttyACM0", &serial.Mode{}))

	defer check.Call(port.Close)

	resetIntoDownloadMode(port)

	time.Sleep(100 * time.Millisecond)

	resetIntoBootingFromFlash(port)
}

// according to "Table 32.4-1. Reset SoC into Download Mode" of "ESP32-C6 Technical Reference Manual"
func resetIntoDownloadMode(port serial.Port) {
	check.Err(port.SetDTR(false))
	check.Err(port.SetRTS(false))
	check.Err(port.SetDTR(true))
	check.Err(port.SetRTS(false))
	check.Err(port.SetRTS(true))
	check.Err(port.SetDTR(false))
	check.Err(port.SetRTS(true))
	check.Err(port.SetRTS(false))
}

// according to "Table 32.4-2. Reset SoC into Booting from flash" of "ESP32-C6 Technical Reference Manual"
func resetIntoBootingFromFlash(port serial.Port) {
	check.Err(port.SetDTR(false))
	check.Err(port.SetRTS(false))
	check.Err(port.SetRTS(true))
	check.Err(port.SetRTS(false))
}
