package dmi

const (
	Data0       = 0x04
	Dmcontrol   = 0x10
	Dmstatus    = 0x11
	Hawindowsel = 0x14
	Hawindow    = 0x15
	Abstractcs  = 0x16
	Command     = 0x17

	DmstatusVersion         = 0
	DmstatusConfstrptrvalid = 4
	DmstatusHasresethaltreq = 5
	DmstatusAuthbusy        = 6
	DmstatusAuthenticated   = 7
	DmstatusAnyhalted       = 8
	DmstatusAllhalted       = 9
	DmstatusAnyrunning      = 10
	DmstatusAllrunning      = 11
	DmstatusAnyunavail      = 12
	DmstatusAllunavail      = 13
	DmstatusAnynonexistent  = 14
	DmstatusAllnonexistent  = 15
	DmstatusAnyresumeack    = 16
	DmstatusAllresumeack    = 17
	DmstatusAnyhavereset    = 18
	DmstatusAllhavereset    = 19
	DmstatusImpebreak       = 22
	DmstatusStickyunavail   = 23
	DmstatusNdmresetpending = 24

	DmcontrolDmactive        = 0
	DmcontrolNdmreset        = 1
	DmcontrolClrresethaltreq = 2
	DmcontrolSetresethaltreq = 3
	DmcontrolClrkeepalive    = 4
	DmcontrolSetkeepalive    = 5
	DmcontrolHartselhi       = 6
	DmcontrolHartsello       = 16
	DmcontrolHasel           = 26
	DmcontrolAckunavail      = 27
	DmcontrolAckhavereset    = 28
	DmcontrolHartreset       = 29
	DmcontrolResumereq       = 30
	DmcontrolHaltreq         = 31

	AbstractcsDatacount   = 0
	AbstractcsCmderr      = 8
	AbstractcsRelaxedpriv = 11
	AbstractcsBusy        = 12
	AbstractcsProgbufsize = 24

	CmderrNone         = 0
	CmderrBusy         = 1
	CmderrNotSupported = 2
	CmderrException    = 3
	CmderrHaltResume   = 4
	CmderrBus          = 5
	CmderrReserver     = 6
	CmderrOther        = 7

	CommandArRegno            = 0
	CommandArWrite            = 16
	CommandArTransfer         = 17
	CommandArPostexec         = 18
	CommandArAarpostincrement = 19
	CommandArAarsize          = 20

	CommandCmdtype = 24

	CmdtypeAccessRegister = 0
	CmdtypeQuickAccess    = 1
	CmdtypeAccessMemory   = 2
)
