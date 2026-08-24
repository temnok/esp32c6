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

func (c *Conn) checkAndWaitCommand() {
	a := c.dmi.Read(dmi.Abstractcs)
	for ; a>>dmi.AbstractcsBusy&1 == 1; a = c.dmi.Read(dmi.Abstractcs) {
	}

	if cmderr := a >> dmi.AbstractcsCmderr & 7; cmderr != dmi.CmderrNone {
		panic(fmt.Errorf("abstract command error: %v", cmderr))
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
		2<<dmi.CommandAarsize|
		1<<dmi.CommandTransfer|
		i<<dmi.CommandRegno)

	c.checkAndWaitCommand()
	return c.dmi.Read(dmi.Data0)
}

func (c *Conn) WriteGPR(i, val int) {
	c.WriteCSR(0x1000+i, val)
}

func (c *Conn) WriteCSR(i, val int) {
	c.dmi.Write(dmi.Data0, val)

	c.dmi.Write(dmi.Command, dmi.CmdtypeAccessRegister<<dmi.CommandCmdtype|
		1<<dmi.CommandWrite|
		2<<dmi.CommandAarsize|
		1<<dmi.CommandTransfer|
		i<<dmi.CommandRegno)

	c.checkAndWaitCommand()
}

func (c *Conn) DataRegisterCount() int {
	return c.dmi.Read(dmi.Abstractcs) >> dmi.AbstractcsDatacount & 15
}

func (c *Conn) waitSbcs() int {
	for {
		if val := c.dmi.Read(dmi.Sbcs); val&(1<<dmi.SbcsSbbusy) == 0 {
			if err := val >> dmi.SbcsSberror & 7; err != 0 {
				panic(fmt.Errorf("bus access error: %v", err))
			}

			return val
		}
	}
}

func (c *Conn) ReadMem(addr int, mem []byte) {
	c.dmi.Write(dmi.Sbcs, 2<<dmi.SbcsSbaccess|
		1<<dmi.SbcsSbautoincrement|
		1<<dmi.SbcsSbreadonaddr|
		1<<dmi.SbcsSbreadondata)

	c.dmi.Write(dmi.Sbaddress0, addr&^3)

	for i := -(addr & 3); i < len(mem); i += 4 {
		c.waitSbcs()
		val := c.dmi.Read(dmi.Sbdata0)

		if i >= 0 {
			mem[i] = byte(val)
		}

		if i+1 >= 0 && i+1 < len(mem) {
			mem[i+1] = byte(val >> 8)
		}

		if i+2 >= 0 && i+2 < len(mem) {
			mem[i+2] = byte(val >> 16)
		}

		if i+3 < len(mem) {
			mem[i+3] = byte(val >> 24)
		}
	}
}

func (c *Conn) WriteMem(addr int, mem []byte) {
	c.dmi.Write(dmi.Sbcs, 2<<dmi.SbcsSbaccess|
		1<<dmi.SbcsSbautoincrement)

	c.dmi.Write(dmi.Sbaddress0, addr&^3)

	for i := -(addr & 3); i < len(mem); i += 4 {
		val := 0

		if i >= 0 {
			val = int(mem[i])
		}

		if i+1 >= 0 && i+1 < len(mem) {
			val |= int(mem[i+1]) << 8
		}

		if i+2 >= 0 && i+2 < len(mem) {
			val |= int(mem[i+2]) << 16
		}

		if i+3 < len(mem) {
			val |= int(mem[i+3]) << 24
		}

		c.dmi.Write(dmi.Sbdata0, val)
		c.waitSbcs()
	}
}
