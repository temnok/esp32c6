package debug

import (
	"fmt"
	"github.com/temnok/esp32c6/dmi"
	"github.com/temnok/esp32c6/isa/csr"
)

type Conn struct {
	dmi *dmi.Conn
}

func (c *Conn) HartCount() int {
	c.dmi.Write(dmi.Dmcontrol, 0x3FF<<dmi.DmcontrolHartsello|1<<dmi.DmcontrolDmactive)
	hartselMax := c.dmi.Read(dmi.Dmcontrol) >> dmi.DmcontrolHartsello & 0x3FF

	for i := 0; i <= hartselMax; i++ {
		c.HartSelect(i)

		if c.dmi.Read(dmi.Dmstatus)>>dmi.DmstatusAnynonexistent&1 == 1 {
			return i
		}
	}

	return hartselMax
}

func (c *Conn) HartSelect(i int) {
	c.dmi.Write(dmi.Dmcontrol, i<<dmi.DmcontrolHartsello|1<<dmi.DmcontrolDmactive)
}

func (c *Conn) HartResetAndHalt(i int) {
	c.dmi.Write(dmi.Dmcontrol, i<<dmi.DmcontrolHartsello|
		1<<dmi.DmcontrolAckhavereset|
		1<<dmi.DmcontrolDmactive)

	c.dmi.Write(dmi.Dmcontrol, i<<dmi.DmcontrolHartsello|
		1<<dmi.DmcontrolSetresethaltreq|
		1<<dmi.DmcontrolDmactive)

	c.dmi.Write(dmi.Dmcontrol, i<<dmi.DmcontrolHartsello|
		1<<dmi.DmcontrolHartreset|
		1<<dmi.DmcontrolDmactive)

	c.dmi.Write(dmi.Dmcontrol, i<<dmi.DmcontrolHartsello|
		1<<dmi.DmcontrolDmactive)

	for c.dmi.Read(dmi.Dmstatus)>>dmi.DmstatusAnyhalted&1 == 0 {
	}
}

func (c *Conn) HartHalt(i int) {
	c.dmi.Write(dmi.Dmcontrol, i<<dmi.DmcontrolHartsello|
		1<<dmi.DmcontrolHaltreq|
		1<<dmi.DmcontrolDmactive)

	c.dmi.Write(dmi.Dmcontrol, i<<dmi.DmcontrolHartsello|
		1<<dmi.DmcontrolDmactive)

	for c.dmi.Read(dmi.Dmstatus)>>dmi.DmstatusAnyhalted&1 == 0 {
	}
}

func (c *Conn) ReadPC() int {
	return c.ReadCSR(csr.Dpc)
}

func (c *Conn) ReadGPR(i int) int {
	return c.ReadCSR(0x1000 + i)
}

func (c *Conn) ReadCSR(i int) int {
	c.dmi.Write(dmi.Command, dmi.CmdtypeAccessRegister<<dmi.CommandCmdtype|
		2<<dmi.CommandARAarsize|
		1<<dmi.CommandARTransfer|
		i<<dmi.CommandARRegno)

	a := c.dmi.Read(dmi.Abstractcs)
	for ; a>>dmi.AbstractcsBusy&1 == 1; a = c.dmi.Read(dmi.Abstractcs) {
	}

	if cmderr := a >> dmi.AbstractcsCmderr & 7; cmderr != dmi.CmderrNone {
		panic(fmt.Errorf("abstract command error: %v", cmderr))
	}

	return c.dmi.Read(dmi.Data0)
}
