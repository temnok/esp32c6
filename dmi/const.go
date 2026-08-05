package dmi

const (
	Data0       = 0x04
	Dmcontrol   = 0x10
	Dmstatus    = 0x11
	Hawindowsel = 0x14
	Hawindow    = 0x15
	Command     = 0x17

	DmcontrolDmactive       = 0
	DmcontrolNdmreset       = 1
	DmcontrolHasel          = 26
	DmstatusNdmresetpending = 24
)
