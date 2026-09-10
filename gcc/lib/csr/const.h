#pragma once

const int
	CSR_USTATUS = 0x000, // User mode STATUS
	CSR_UIE     = 0x004, // User Interrupt Enable
	CSR_UTVEC   = 0x005, // User Trap VECtor

	CSR_USCRATCH = 0x040, // User SCRATCH register
	CSR_UEPC     = 0x041, // User Exception Program Counter
	CSR_UCAUSE   = 0x042, // User trap CAUSE
	CSR_UIP      = 0x044, // User Interrupt Pending

	CSR_MSTATUS = 0x300, // Machine mode STATUS
	CSR_MISA    = 0x301, // Machine ISA
	CSR_MIDELEG = 0x303, // Machine Interrupt DELEGation register
	CSR_MIE     = 0x304, // Machine Interrupt Enable
	CSR_MTVEC   = 0x305, // Machine Trap VECtor

	CSR_MSCRATCH = 0x340, // Machine SCRATCH register
	CSR_MEPC     = 0x341, // Machine Exception Program Counter
	CSR_MCAUSE   = 0x342, // Machine trap CAUSE
	CSR_MTVAL    = 0x343, // Machine Trap VALue
	CSR_MIP      = 0x344, // Machine Interrupt Pending

	CSR_PMPCFG0 = 0x3A0, // Physical Memory Protection ConFiGuration
	CSR_PMPCFG1 = 0x3A1,
	CSR_PMPCFG2 = 0x3A2,
	CSR_PMPCFG3 = 0x3A3,

	CSR_PMPADDR0  = 0x3B0, // Physical Memory Protection ADDRess
	CSR_PMPADDR1  = 0x3B1,
	CSR_PMPADDR2  = 0x3B2,
	CSR_PMPADDR3  = 0x3B3,
	CSR_PMPADDR4  = 0x3B4,
	CSR_PMPADDR5  = 0x3B5,
	CSR_PMPADDR6  = 0x3B6,
	CSR_PMPADDR7  = 0x3B7,
	CSR_PMPADDR8  = 0x3B8,
	CSR_PMPADDR9  = 0x3B9,
	CSR_PMPADDR10 = 0x3BA,
	CSR_PMPADDR11 = 0x3BB,
	CSR_PMPADDR12 = 0x3BC,
	CSR_PMPADDR13 = 0x3BD,
	CSR_PMPADDR14 = 0x3BE,
	CSR_PMPADDR15 = 0x3BF,

	CSR_TSELECT  = 0x7A0, // Trigger SELECT register
	CSR_TDATA1   = 0x7A1, // Trigger abstract DATA 1
	CSR_TDATA2   = 0x7A2, // Trigger abstract DATA 1
	CSR_TCONTROL = 0x7A5, // Trigger CONTROL register

	CSR_DCSR      = 0x7B0, // Debug Control and Status Register
	CSR_DPC       = 0x7B1, // Debug PC
	CSR_DSCRATCH0 = 0x7B2, // Debug SCRATCH Register 0
	CSR_DSCRATCH1 = 0x7B3, // Debug SCRATCH Register 1

	CSR_MPCER = 0x7E0, // Machine Performance Counter Event Register
	CSR_MPCMR = 0x7E1, // Machine Performance Counter Mode Register
	CSR_MPCCR = 0x7E2, // Machine Performance Counter Count Register

	CSR_CPU_GPIO_OEN = 0x803, // CPU GPIO Output ENable
	CSR_CPU_GPIO_IN  = 0x804, // CPU GPIO INput value
	CSR_CPU_GPIO_OUT = 0x805, // CPU GPIO OUTput value

	CSR_PMA_CFG0  = 0xBC0, // Physical Memory Attribute ConFiGuration
	CSR_PMA_CFG1  = 0xBC1,
	CSR_PMA_CFG2  = 0xBC2,
	CSR_PMA_CFG3  = 0xBC3,
	CSR_PMA_CFG4  = 0xBC4,
	CSR_PMA_CFG5  = 0xBC5,
	CSR_PMA_CFG6  = 0xBC6,
	CSR_PMA_CFG7  = 0xBC7,
	CSR_PMA_CFG8  = 0xBC8,
	CSR_PMA_CFG9  = 0xBC9,
	CSR_PMA_CFG10 = 0xBCA,
	CSR_PMA_CFG11 = 0xBCB,
	CSR_PMA_CFG12 = 0xBCC,
	CSR_PMA_CFG13 = 0xBCD,
	CSR_PMA_CFG14 = 0xBCE,
	CSR_PMA_CFG15 = 0xBCF,

	CSR_PMA_ADDR0  = 0xBD0,
	CSR_PMA_ADDR1  = 0xBD1,
	CSR_PMA_ADDR2  = 0xBD2,
	CSR_PMA_ADDR3  = 0xBD3,
	CSR_PMA_ADDR4  = 0xBD4,
	CSR_PMA_ADDR5  = 0xBD5,
	CSR_PMA_ADDR6  = 0xBD6,
	CSR_PMA_ADDR7  = 0xBD7,
	CSR_PMA_ADDR8  = 0xBD8,
	CSR_PMA_ADDR9  = 0xBD9,
	CSR_PMA_ADDR10 = 0xBDA,
	CSR_PMA_ADDR11 = 0xBDB,
	CSR_PMA_ADDR12 = 0xBDC,
	CSR_PMA_ADDR13 = 0xBDD,
	CSR_PMA_ADDR14 = 0xBDE,

	CSR_MVENDORID = 0xF11, // Machine VENDOR ID
	CSR_MARCHID   = 0xF12, // Machine ARCHitecture ID
	CSR_MIMPID    = 0xF13, // Machine IMPlementation ID
	CSR_MHARTID   = 0xF14, // Machine HART ID

	CSR_MSTATUS_UIE = 0,
	CSR_MSTATUS_MIE = 1,
	CSR_MSTATUS_MPP = 11,
	CSR_MSTATUS_TW  = 21,

	CSR_DCSR_PRV     = 0,
	CSR_DCSR_CAUSE   = 6,
	CSR_DCSR_EBREAKU = 12,
	CSR_DCSR_EBREAKM = 15,

	CSR_MPCER_CYCLE = 0;
