package isa

type RV32S interface {
	MRET()                   // Machine mode RETurn
	SFENCE_VMA(rs1, rs2 int) // Supervisor FENCE for Virtual Memory Address
	SRET()                   // Supervisor mode RETurn
	URET()                   // User mode RETurn
	WFI()                    // Wait For Interrupt
}
