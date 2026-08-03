package tests

import (
	"encoding/binary"
	"github.com/google/gousb"
	"github.com/stretchr/testify/assert"
	"github.com/temnok/esp32c6/check"
	"testing"
)

func TestTapIdcode(t *testing.T) {
	defer handlePanic(t)

	withJtagUsbEndpoints(func(in *gousb.InEndpoint, out *gousb.OutEndpoint) {

		sequence := []byte{
			0b_0010_0010, // send 5 TMS clocks to reset state machine (same as TRST)
			0b_0010_0010, // entering TestLogicReset state
			0b_0010_0000, // and then RunTestIdle state; by default, IR == IDCODE

			0b_0010_0000, // SelectDRScan, CaptureDR

			0b_0000_0100, // 16x2 = 32 ShiftDR with capturing TDI
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

			0b_0110_0010, // Exit1DR, UpdateDR
			0b_0000_1010, // RunTestIdle, Flush
		}

		assert.Equal(t, 22, check.Err1(out.Write(sequence)))

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
