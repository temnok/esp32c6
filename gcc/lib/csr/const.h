#pragma once

const int
	csr_ustatus = 0x000, // User mode STATUS
	csr_uie     = 0x004, // User Interrupt Enable
	csr_utvec   = 0x005, // User Trap VECtor

	csr_uscratch = 0x040, // User SCRATCH register
	csr_uepc     = 0x041, // User Exception Program Counter
	csr_ucause   = 0x042, // User trap CAUSE
	csr_uip      = 0x044, // User Interrupt Pending

	csr_mstatus = 0x300, // Machine mode STATUS
	csr_misa    = 0x301, // Machine ISA
	csr_mideleg = 0x303, // Machine Interrupt DELEGation register
	csr_mie     = 0x304, // Machine Interrupt Enable
	csr_mtvec   = 0x305, // Machine Trap VECtor

	csr_mscratch = 0x340, // Machine SCRATCH register
	csr_mepc     = 0x341, // Machine Exception Program Counter
	csr_mcause   = 0x342, // Machine trap CAUSE
	csr_mtval    = 0x343, // Machine Trap VALue
	csr_mip      = 0x344, // Machine Interrupt Pending

	csr_pmpcfg0 = 0x3A0, // Physical Memory Protection ConFiGuration
	csr_pmpcfg1 = 0x3A1,
	csr_pmpcfg2 = 0x3A2,
	csr_pmpcfg3 = 0x3A3,

	csr_pmpaddr0  = 0x3B0, // Physical Memory Protection ADDRess
	csr_pmpaddr1  = 0x3B1,
	csr_pmpaddr2  = 0x3B2,
	csr_pmpaddr3  = 0x3B3,
	csr_pmpaddr4  = 0x3B4,
	csr_pmpaddr5  = 0x3B5,
	csr_pmpaddr6  = 0x3B6,
	csr_pmpaddr7  = 0x3B7,
	csr_pmpaddr8  = 0x3B8,
	csr_pmpaddr9  = 0x3B9,
	csr_pmpaddr10 = 0x3BA,
	csr_pmpaddr11 = 0x3BB,
	csr_pmpaddr12 = 0x3BC,
	csr_pmpaddr13 = 0x3BD,
	csr_pmpaddr14 = 0x3BE,
	csr_pmpaddr15 = 0x3BF,

	csr_tselect  = 0x7A0, // Trigger SELECT register
	csr_tdata1   = 0x7A1, // Trigger abstract DATA 1
	csr_tdata2   = 0x7A2, // Trigger abstract DATA 1
	csr_tcontrol = 0x7A5, // Trigger CONTROL register

	csr_dcsr      = 0x7B0, // Debug Control and Status Register
	csr_dpc       = 0x7B1, // Debug PC
	csr_dscratch0 = 0x7B2, // Debug SCRATCH Register 0
	csr_dscratch1 = 0x7B3, // Debug SCRATCH Register 1

	csr_mpcer = 0x7E0, // Machine Performance Counter Event Register
	csr_mpcmr = 0x7E1, // Machine Performance Counter Mode Register
	csr_mpccr = 0x7E2, // Machine Performance Counter Count Register

	csr_cpu_gpio_oen = 0x803, // CPU GPIO Output ENable
	csr_cpu_gpio_in  = 0x804, // CPU GPIO INput value
	csr_cpu_gpio_out = 0x805, // CPU GPIO OUTput value

	csr_pma_cfg0  = 0xBC0, // Physical Memory Attribute ConFiGuration
	csr_pma_cfg1  = 0xBC1,
	csr_pma_cfg2  = 0xBC2,
	csr_pma_cfg3  = 0xBC3,
	csr_pma_cfg4  = 0xBC4,
	csr_pma_cfg5  = 0xBC5,
	csr_pma_cfg6  = 0xBC6,
	csr_pma_cfg7  = 0xBC7,
	csr_pma_cfg8  = 0xBC8,
	csr_pma_cfg9  = 0xBC9,
	csr_pma_cfg10 = 0xBCA,
	csr_pma_cfg11 = 0xBCB,
	csr_pma_cfg12 = 0xBCC,
	csr_pma_cfg13 = 0xBCD,
	csr_pma_cfg14 = 0xBCE,
	csr_pma_cfg15 = 0xBCF,

	csr_pma_addr0  = 0xBD0,
	csr_pma_addr1  = 0xBD1,
	csr_pma_addr2  = 0xBD2,
	csr_pma_addr3  = 0xBD3,
	csr_pma_addr4  = 0xBD4,
	csr_pma_addr5  = 0xBD5,
	csr_pma_addr6  = 0xBD6,
	csr_pma_addr7  = 0xBD7,
	csr_pma_addr8  = 0xBD8,
	csr_pma_addr9  = 0xBD9,
	csr_pma_addr10 = 0xBDA,
	csr_pma_addr11 = 0xBDB,
	csr_pma_addr12 = 0xBDC,
	csr_pma_addr13 = 0xBDD,
	csr_pma_addr14 = 0xBDE,

	csr_mvendorid = 0xF11, // Machine VENDOR ID
	csr_marchid   = 0xF12, // Machine ARCHitecture ID
	csr_mimpid    = 0xF13, // Machine IMPlementation ID
	csr_mhartid   = 0xF14, // Machine HART ID

	csr_mstatus_UIE = 0,
	csr_mstatus_MIE = 1,
	csr_mstatus_MPP = 11,
	csr_mstatus_TW  = 21,

	csr_dcsr_PRV     = 0,
	csr_dcsr_CAUSE   = 6,
	csr_dcsr_EBREAKU = 12,
	csr_dcsr_EBREAKM = 15,

	csr_mpcer_CYCLE = 0;
