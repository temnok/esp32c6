package tests

import (
	"github.com/google/gousb"
	"github.com/stretchr/testify/assert"
	"github.com/temnok/esp32c6/check"
	"testing"
)

func TestJtagBypass(t *testing.T) {
	defer handlePanic(t)

	withJtagUsbEndpoints(func(in *gousb.InEndpoint, out *gousb.OutEndpoint) {

		sequence := []byte{
			0b_0010_0010, // Reset, Reset,
			0b_0010_0010, // Reset, Reset,
			0b_0010_0000, // Reset, Idle,
			0b_0010_0010, // Select-DR, Select-IR,
			0b_0000_0000, // Capture-IR, Shift-IR,
			0b_0000_0000, // Shift-IR, Shift-IR,
			0b_0000_0000, // Shift-IR, Shift-IR,
			0b_0010_0010, // Exit1-IR, Update-IR,
			0b_0000_0010, // Idle, Select-DR,
			0b_0000_0001, // Capture-DR, Shift-DR,
			0b_0101_0100, // Shift-DR, Shift-DR,
			0b_0101_0101, // Shift-DR, Shift-DR,
			0b_0100_0101, // Shift-DR, Shift-DR,
			0b_0101_0110, // Shift-DR, Exit1-DR,
			0b_0010_1010, // Update-DR, Flush
		}

		assert.Equal(t, len(sequence), check.Err1(out.Write(sequence)))

		data := make([]byte, 1)
		assert.Equal(t, 1, check.Err1(in.Read(data)))

		assert.Equal(t, 0b_11011010, int(data[0]))
	})
}
