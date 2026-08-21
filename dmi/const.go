package dmi

const (
	Data0       = 0x04
	Data1       = 0x08
	Dmcontrol   = 0x10
	Dmstatus    = 0x11
	Hawindowsel = 0x14
	Hawindow    = 0x15
	Abstractcs  = 0x16
	Command     = 0x17
	Sbcs        = 0x38
	Sbaddress0  = 0x39
	Sbdata0     = 0x3C

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

	CommandRegno    = 0
	CommandWrite    = 16
	CommandTransfer = 17
	CommandPostexec = 18

	CommandAarpostincrement = 19
	CommandAarsize          = 20

	CommandAampostincrement = 19
	CommandAamsize          = 20

	CommandCmdtype = 24

	CmdtypeAccessRegister = 0
	CmdtypeQuickAccess    = 1
	CmdtypeAccessMemory   = 2

	SbcsSbaccess8       = 0
	SbcsSbaccess16      = 1
	SbcsSbaccess32      = 2
	SbcsSbaccess64      = 3
	SbcsSbaccess128     = 4
	SbcsSbasize         = 5
	SbcsSberror         = 12
	SbcsSbreadondata    = 15
	SbcsSbautoincrement = 16
	SbcsSbaccess        = 17
	SbcsSbreadonaddr    = 20
	SbcsSbbusy          = 21
	SbcsSbbusyerror     = 22
	SbcsSbversion       = 29

	SberrorNone      = 0
	SberrorTimout    = 1
	SberrorAddress   = 2
	SberrorAlignment = 3
	SberrorSize      = 4
	SberrorOther     = 7
)
