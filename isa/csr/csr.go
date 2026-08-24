package csr

const (
	Ustatus = 0x000 // User mode STATUS
	Uie     = 0x004 // User Interrupt Enable
	Utvec   = 0x005 // User Trap VECtor

	Uscratch = 0x040 // User SCRATCH register
	Uepc     = 0x041 // User Exception Program Counter
	Ucause   = 0x042 // User trap CAUSE
	Uip      = 0x044 // User Interrupt Pending

	Mstatus = 0x300 // Machine mode STATUS
	Misa    = 0x301 // Machine ISA
	Mideleg = 0x303 // Machine Interrupt DELEGation register
	Mie     = 0x304 // Machine Interrupt Enable
	Mtvec   = 0x305 // Machine Trap VECtor

	Mscratch = 0x340 // Machine SCRATCH register
	Mepc     = 0x341 // Machine Exception Program Counter
	Mcause   = 0x342 // Machine trap CAUSE
	Mtval    = 0x343 // Machine Trap VALue
	Mip      = 0x344 // Machine Interrupt Pending

	Pmpcfg0 = 0x3A0 // Physical Memory Protection ConFiGuration
	Pmpcfg1 = 0x3A1
	Pmpcfg2 = 0x3A2
	Pmpcfg3 = 0x3A3

	Pmpaddr0  = 0x3B0 // Physical Memory Protection ADDRess
	Pmpaddr1  = 0x3B1
	Pmpaddr2  = 0x3B2
	Pmpaddr3  = 0x3B3
	Pmpaddr4  = 0x3B4
	Pmpaddr5  = 0x3B5
	Pmpaddr6  = 0x3B6
	Pmpaddr7  = 0x3B7
	Pmpaddr8  = 0x3B8
	Pmpaddr9  = 0x3B9
	Pmpaddr10 = 0x3BA
	Pmpaddr11 = 0x3BB
	Pmpaddr12 = 0x3BC
	Pmpaddr13 = 0x3BD
	Pmpaddr14 = 0x3BE
	Pmpaddr15 = 0x3BF

	Tselect  = 0x7A0 // Trigger SELECT register
	Tdata1   = 0x7A1 // Trigger abstract DATA 1
	Tdata2   = 0x7A2 // Trigger abstract DATA 1
	Tcontrol = 0x7A5 // Trigger CONTROL register

	Dcsr      = 0x7B0 // Debug Control and Status Register
	Dpc       = 0x7B1 // Debug PC
	Dscratch0 = 0x7B2 // Debug SCRATCH Register 0
	Dscratch1 = 0x7B3 // Debug SCRATCH Register 1

	Mpcer = 0x7E0 // Machine Performance Counter Event Register
	Mpcmr = 0x7E1 // Machine Performance Counter Mode Register
	Mpccr = 0x7E2 // Machine Performance Counter Count Register

	Cpu_gpio_oen = 0x803 // CPU GPIO Output ENable
	Cpu_gpio_in  = 0x804 // CPU GPIO INput value
	Cpu_gpio_out = 0x805 // CPU GPIO OUTput value

	Pma_cfg0  = 0xBC0 // Physical Memory Attribute ConFiGuration
	Pma_cfg1  = 0xBC1
	Pma_cfg2  = 0xBC2
	Pma_cfg3  = 0xBC3
	Pma_cfg4  = 0xBC4
	Pma_cfg5  = 0xBC5
	Pma_cfg6  = 0xBC6
	Pma_cfg7  = 0xBC7
	Pma_cfg8  = 0xBC8
	Pma_cfg9  = 0xBC9
	Pma_cfg10 = 0xBCA
	Pma_cfg11 = 0xBCB
	Pma_cfg12 = 0xBCC
	Pma_cfg13 = 0xBCD
	Pma_cfg14 = 0xBCE
	Pma_cfg15 = 0xBCF

	Pma_addr0  = 0xBD0
	Pma_addr1  = 0xBD1
	Pma_addr2  = 0xBD2
	Pma_addr3  = 0xBD3
	Pma_addr4  = 0xBD4
	Pma_addr5  = 0xBD5
	Pma_addr6  = 0xBD6
	Pma_addr7  = 0xBD7
	Pma_addr8  = 0xBD8
	Pma_addr9  = 0xBD9
	Pma_addr10 = 0xBDA
	Pma_addr11 = 0xBDB
	Pma_addr12 = 0xBDC
	Pma_addr13 = 0xBDD
	Pma_addr14 = 0xBDE

	Mvendorid = 0xF11 // Machine VENDOR ID
	Marchid   = 0xF12 // Machine ARCHitecture ID
	Mimpid    = 0xF13 // Machine IMPlementation ID
	Mhartid   = 0xF14 // Machine HART ID
)

const (
	MstatusUIE = 0
	MstatusMIE = 1
	MstatusMPP = 11
	MstatusTW  = 21

	DcsrPrv     = 0
	DcsrCause   = 6
	DcsrEbreaku = 12
	DcsrEbreakm = 15

	MpcerCycle = 0
)
