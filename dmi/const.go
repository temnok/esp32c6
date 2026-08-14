package dmi

const (
	Data0       = 0x04
	Dmcontrol   = 0x10
	Dmstatus    = 0x11
	Hawindowsel = 0x14
	Hawindow    = 0x15
	Command     = 0x17

	DmcontrolDmactive        = 0
	DmcontrolNdmreset        = 1
	DmcontrolClrresethaltreq = 2
	DmcontrolSetresethaltreq = 3
	DmcontrolHartselhi       = 6
	DmcontrolHartsello       = 16
	DmcontrolHasel           = 26
	DmcontrolAckhavereset    = 28
	DmcontrolHartreset       = 29
	DmcontrolResumereq       = 30
	DmcontrolHaltreq         = 31

	DmstatusAnyhalted       = 8
	DmstatusAllhalted       = 9
	DmstatusAnyrunning      = 10
	DmstatusAnynonexistent  = 14
	DmstatusAllhavereset    = 19
	DmstatusNdmresetpending = 24
)
