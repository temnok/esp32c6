package debug

import (
	"github.com/temnok/esp32c6/dmi"
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

func (c *Conn) HartHalt(i int) {
	c.dmi.Write(dmi.Dmcontrol, i<<dmi.DmcontrolHartsello|1<<dmi.DmcontrolHaltreq|1<<dmi.DmcontrolDmactive)

	for c.dmi.Read(dmi.Dmstatus)>>dmi.DmstatusAnyhalted&1 == 0 {
	}

	c.dmi.Write(dmi.Dmcontrol, i<<dmi.DmcontrolHartsello|1<<dmi.DmcontrolDmactive)
}
