package dmi

import (
	"fmt"
	"github.com/temnok/esp32c6/jtag"
)

func initialize(conn *jtag.Conn) {
	conn.WriteIR(0x11)

	write(conn, Dmcontrol, 0<<DmcontrolDmactive)
	for read(conn, Dmcontrol)>>DmcontrolDmactive&1 != 0 {
	}

	write(conn, Dmcontrol, 1<<DmcontrolDmactive)
	for read(conn, Dmcontrol)>>DmcontrolDmactive&1 == 0 {
	}
}

// Debug Module Interface Access (dmi, at 0x11)
func access(conn *jtag.Conn, addr, data, op int) (int, int, int) {
	a, d := addr, data
	nibbles := []byte{
		0x20,                                    // Select-DR, Capture-DR,
		0x00 | byte(op&1),                       // Shift-DR, Shift-DR,
		0x00 | byte(op>>1&1)<<4 | byte(d&1),     // Shift-DR, Shift-DR,
		0x00 | byte(d>>1&1)<<4 | byte(d>>2&1),   // Shift-DR, Shift-DR,
		0x00 | byte(d>>3&1)<<4 | byte(d>>4&1),   // Shift-DR, Shift-DR,
		0x00 | byte(d>>5&1)<<4 | byte(d>>6&1),   // Shift-DR, Shift-DR,
		0x00 | byte(d>>7&1)<<4 | byte(d>>8&1),   // Shift-DR, Shift-DR,
		0x00 | byte(d>>9&1)<<4 | byte(d>>10&1),  // Shift-DR, Shift-DR,
		0x00 | byte(d>>11&1)<<4 | byte(d>>12&1), // Shift-DR, Shift-DR,
		0x00 | byte(d>>13&1)<<4 | byte(d>>14&1), // Shift-DR, Shift-DR,
		0x00 | byte(d>>15&1)<<4 | byte(d>>16&1), // Shift-DR, Shift-DR,
		0x00 | byte(d>>17&1)<<4 | byte(d>>18&1), // Shift-DR, Shift-DR,
		0x00 | byte(d>>19&1)<<4 | byte(d>>20&1), // Shift-DR, Shift-DR,
		0x00 | byte(d>>21&1)<<4 | byte(d>>22&1), // Shift-DR, Shift-DR,
		0x00 | byte(d>>23&1)<<4 | byte(d>>24&1), // Shift-DR, Shift-DR,
		0x00 | byte(d>>25&1)<<4 | byte(d>>26&1), // Shift-DR, Shift-DR,
		0x00 | byte(d>>27&1)<<4 | byte(d>>28&1), // Shift-DR, Shift-DR,
		0x00 | byte(d>>29&1)<<4 | byte(d>>30&1), // Shift-DR, Shift-DR,
		0x00 | byte(d>>31&1)<<4 | byte(a&1),     // Shift-DR, Shift-DR,
		0x00 | byte(a>>1&1)<<4 | byte(a>>2&1),   // Shift-DR, Shift-DR,
		0x00 | byte(a>>3&1)<<4 | byte(a>>4&1),   // Shift-DR, Shift-DR,
		0x02 | byte(a>>5&1)<<4 | byte(a>>6&1),   // Shift-DR, Exit1-DR,
		0x20,                                    // Update-DR, Idle,
		0x00,                                    // Idle, Idle,
		0x02,                                    // Idle, Select-DR,
		0x00,                                    // Capture-DR, Shift-DR,
		0x4F,                                    // Shift-DR with capture, 7 x Shift-DR with capture
		0xDE,                                    // 32 x Shift-DR with capture,
		0x62,                                    // Exit1-DR with capture, Update-DR,
		0x0A,                                    // Idle, Flush
	}

	conn.Write(nibbles)

	r := make([]byte, 6)
	conn.Read(r)

	resp := int(r[5])<<40 | int(r[4])<<32 | int(r[3])<<24 | int(r[2])<<16 | int(r[1])<<8 | int(r[0])
	return resp >> 34, int(uint32(resp >> 2)), resp & 3
}

func read(conn *jtag.Conn, addr int) int {
	a, d, op := access(conn, addr, 0, 1)
	if op != 0 {
		panic(fmt.Errorf("ReadDMI returned %v", op))
	}

	if a != addr {
		panic(fmt.Errorf("ReadDMI returned address 0x%X, expected 0x%X", a, addr))
	}

	return d
}

func write(conn *jtag.Conn, addr, data int) {
	a, _, op := access(conn, addr, data, 2)
	if op != 0 {
		panic(fmt.Errorf("WriteDMI returned %v", op))
	}

	if a != addr {
		panic(fmt.Errorf("WriteDMI returned address 0x%X, expected 0x%X", a, addr))
	}
}
