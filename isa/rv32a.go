package isa

type RV32A interface {
	AMOADD_W(rd, rs2, rs1 int)
	AMOAND_W(rd, rs2, rs1 int)
	AMOMAXU_W(rd, rs2, rs1 int)
	AMOMAX_W(rd, rs2, rs1 int)
	AMOMINU_W(rd, rs2, rs1 int)
	AMOMIN_W(rd, rs2, rs1 int)
	AMOOR_W(rd, rs2, rs1 int)
	AMOSWAP_W(rd, rs2, rs1 int)
	AMOXOR_W(rd, rs2, rs1 int)
	LR_W(rd, rs1 int)
	SC_W(rd, rs2, rs1 int)
}
