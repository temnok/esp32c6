package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/temnok/esp32c6/check"
	"github.com/temnok/esp32c6/dmi"
	"testing"
)

func TestDmiDmstatus(t *testing.T) {
	defer check.RecoverAndPrintStack(t.Fatal)

	dmi.Session(func(conn *dmi.Conn) {
		assert.Equal(t, 2, conn.Read(dmi.Dmstatus)&0xF)

		// version         0010 (0.13): There is a Debug Module and it conforms to version 0.13 of this specification
		// confstrptrvalid    0 (invalid): confstrptr0--confstrptr3 hold information which is not relevant to the configuration structure
		// hasresethaltreq    1 if this Debug Module supports halt-on-reset functionality controllable by the setresethaltreq and clrresethaltreq bits
		// authbusy           0 (ready): The authentication module is ready to process the next read/write to authdata
		// authenticated      1 (true): The authentication check has passed
		// anyhalted          0 This field is 1 when any currently selected hart is halted
		// allhalted          0 This field is 1 when all currently selected harts are halted
		// anyrunning         1 This field is 1 when any currently selected hart is running
		// allrunning         1 This field is 1 when all currently selected harts are running
		// anyunavail         0 This field is 1 when any currently selected hart is unavailable
		// allunavail         0 This field is 1 when all currently selected harts are unavailable
		// anynonexistent     0 This field is 1 when any currently selected hart does not exist in this hardware platform
		// allnonexistent     0 This field is 1 when all currently selected harts do not exist in this hardware platform
		// anyresumeack       0 This field is 1 when any currently selected hart has its resume ack bit set
		// allresumeack       0 This field is 1 when all currently selected harts have their resume ack bit set
		// anyhavereset       1 This field is 1 when at least one currently selected hart has been reset and reset has not been acknowledged for that hart
		// allhavereset       1 This field is 1 when all currently selected harts have been reset and reset has not been acknowledged for any of them
	})
}

func TestDmcontrolHasel(t *testing.T) {
	defer check.RecoverAndPrintStack(t.Fatal)

	dmi.Session(func(conn *dmi.Conn) {
		conn.Write(dmi.Dmcontrol, 1<<dmi.DmcontrolHasel|1<<dmi.DmcontrolDmactive)
		assert.Equal(t, 1, conn.Read(dmi.Dmcontrol)>>dmi.DmcontrolHasel&1)

		conn.Write(dmi.Dmcontrol, 0<<dmi.DmcontrolHasel|1<<dmi.DmcontrolDmactive)
		assert.Equal(t, 0, conn.Read(dmi.Dmcontrol)>>dmi.DmcontrolHasel&1)
	})
}

func TestDmcontrolNdmreset(t *testing.T) {
	defer check.RecoverAndPrintStack(t.Fatal)

	dmi.Session(func(conn *dmi.Conn) {
		conn.Write(dmi.Dmcontrol, 1<<dmi.DmcontrolNdmreset|1<<dmi.DmcontrolDmactive)
		assert.Equal(t, 1, conn.Read(dmi.Dmcontrol)>>dmi.DmcontrolNdmreset&1)

		conn.Write(dmi.Dmcontrol, 0<<dmi.DmcontrolNdmreset|1<<dmi.DmcontrolDmactive)
		assert.Equal(t, 0, conn.Read(dmi.Dmcontrol)>>dmi.DmcontrolNdmreset&1)
	})
}
