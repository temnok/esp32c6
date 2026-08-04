package tests

import (
	"encoding/binary"
	"github.com/google/gousb"
	"github.com/stretchr/testify/assert"
	"github.com/temnok/esp32c6/check"
	"testing"
)

func TestJtagIdcode(t *testing.T) {
	defer handlePanic(t)

	withJtagUsbEndpoints(func(in *gousb.InEndpoint, out *gousb.OutEndpoint) {

		// Table 32.3-3. Commands of a Nibble

		// CMD_CLK   0 cap tms tdi
		// CMD_RST   1   0   0 rst
		// CMD_FLUSH 1   0   1   0
		// CMD_RSV   1   0   1   1
		// CMD_REP   1   1  R1  R0

		sequence := []byte{
			0b_0010_0010, // send 5 TMS clocks to reset state machine (same as TRST)
			0b_0010_0010, // entering Reset state
			0b_0010_0000, // and then Idle state; by default, IR == IDCODE

			0b_0010_0000, // Select-DR, Capture-DR

			0b_0000_0100, // 16x2 = 32 Shift-DR with capturing TDI
			0b_0100_0100,
			0b_0100_0100,
			0b_0100_0100,
			0b_0100_0100,
			0b_0100_0100,
			0b_0100_0100,
			0b_0100_0100,
			0b_0100_0100,
			0b_0100_0100,
			0b_0100_0100,
			0b_0100_0100,
			0b_0100_0100,
			0b_0100_0100,
			0b_0100_0100,
			0b_0100_0100,

			0b_0110_0010, // Exit1-DR, Update-DR
			0b_0000_1010, // Idle, Flush
		}

		assert.Equal(t, len(sequence), check.Err1(out.Write(sequence)))

		data := make([]byte, 4)
		assert.Equal(t, 4, check.Err1(in.Read(data)))

		// https://docs.espressif.com/projects/esp-idf/en/stable/esp32c6/api-guides/jtag-debugging/index.html#codecell3
		assert.Equal(t, []byte{0x25, 0xDC, 0x0, 0x0}, data)

		idcode := binary.LittleEndian.Uint32(data)
		assert.Equal(t, uint32(0xDC25), idcode)

		version := idcode >> 28
		partNumber := idcode >> 12 & 0xFFFF
		manufID := idcode >> 1 & 0x7FF

		assert.Equal(t, uint32(0), version)
		assert.Equal(t, uint32(0xD), partNumber) // ESP32-C6
		assert.Equal(t, uint32(0x612), manufID)  // Espressif Systems

	})
}
