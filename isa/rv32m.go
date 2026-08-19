package isa

type RV32M interface {
	DIV(rd, rs1, rs2 int)    // DIVide
	DIVU(rd, rs1, rs2 int)   // DIVide Unsigned
	MUL(rd, rs1, rs2 int)    // MULtiply
	MULH(rd, rs1, rs2 int)   // MULtiply High
	MULHSU(rd, rs1, rs2 int) // MULtiply High Signed and Unsigned
	MULHU(rd, rs1, rs2 int)  // MULtiply High Unsigned
	REM(rd, rs1, rs2 int)    // REMinder
	REMU(rd, rs1, rs2 int)   // REMinder of Unsigned
}
