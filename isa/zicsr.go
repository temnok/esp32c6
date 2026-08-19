package isa

type Zicsr interface {
	CSRRC(rd, csr, rs1 int)
	CSRRCI(rd, csr, imm int)
	CSRRS(rd, csr, rs1 int)
	CSRRSI(rd, csr, imm int)
	CSRRW(rd, csr, rs1 int)
	CSRRWI(rd, csr, imm int)
}
