package isa

type Privileged interface {
	MRET()                   // Machine mode RETurn
	SFENCE_VMA(rs1, rs2 int) // Supervisor FENCE for Virtual Memory Address
	SRET()                   // Supervisor mode RETurn
	URET()                   // User mode RETurn
	WFI()                    // Wait For Interrupt
}

const (
	MRET       = 0x30200073
	SFENCE_VMA = 0x12000073
	SRET       = 0x10200073
	URET       = 0x00200073
	WFI        = 0x10500073
)
