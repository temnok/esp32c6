package isa

type RV32I interface { //               42 =
	baseComputationalRegInstructions // 10 +
	baseComputationalImmInstructions // 11 +
	baseControlTransferInstructions  //  8 +
	baseLoadStoreInstructions        //  8 +
	baseMemoryOrderingInstructions   //  1 +
	baseSystemInstructions           //  4
}

type baseComputationalRegInstructions interface {
	ADD(rd, rs1, rs2 int)  // ADD
	AND(rd, rs1, rs2 int)  // AND
	OR(rd, rs1, rs2 int)   // OR
	SLL(rd, rs1, rs2 int)  // Shift Left, Logical
	SLT(rd, rs1, rs2 int)  // Set if Less Than
	SLTU(rd, rs1, rs2 int) // Set if Less Than, Unsigned
	SRA(rd, rs1, rs2 int)  // Shift Right, Arithmetical
	SRL(rd, rs1, rs2 int)  // Shift Right, Logical
	SUB(rd, rs1, rs2 int)  // SUBtract
	XOR(rd, rs1, rs2 int)  // eXclusive OR
}

type baseComputationalImmInstructions interface {
	ADDI(rd, rs1, imm int)  // ADD Immediate
	ANDI(rd, rs1, imm int)  // AND with Immediate
	AUIPC(rd, imm int)      // Add Upper Immediate to PC
	LUI(rd, imm int)        // Load Upper Immediate
	ORI(rd, rs1, imm int)   // OR with Immediate
	SLLI(rd, rs1, imm int)  // Shift Left, Logical by Immediate
	SLTI(rd, rs1, imm int)  // Set if Less Than, with Immediate
	SLTIU(rd, rs1, imm int) // Set if Less Than, with Immediate, Unsigned
	SRAI(rd, rs1, imm int)  // Shift Right, Arithmetical by Immediate
	SRLI(rd, rs1, imm int)  // Shift Right, Logical with Immediate
	XORI(rd, rs1, imm int)  // eXclusive OR with Immediate
}

type baseControlTransferInstructions interface {
	BEQ(rs1, rs2, imm int)  // Branch if EQual
	BGE(rs1, rs2, imm int)  // Branch if Greater or Equal
	BGEU(rs1, rs2, imm int) // Branch if Greater or Equal, Unsigned operands
	BLT(rs1, rs2, imm int)  // Branch if Less Than
	BLTU(rs1, rs2, imm int) // Branch if Less Than, Unsigned operands
	BNE(rs1, rs2, imm int)  // Branch if Not Equal
	JAL(rd, imm int)        // Jump And Link by immediate offset
	JALR(rd, rs1, imm int)  // Jump And Link to address in register
}

type baseLoadStoreInstructions interface {
	LB(rd, rs1, imm int)  // Load Byte
	LBU(rd, rs1, imm int) // Load Byte, Unsigned
	LH(rd, rs1, imm int)  // Load Half word
	LHU(rd, rs1, imm int) // Load Half word, Unsigned
	LW(rd, rs1, imm int)  // Load Word
	SB(rs2, rs1, imm int) // Store Byte
	SH(rs2, rs1, imm int) // Store Half word
	SW(rs2, rs1, imm int) // Store Word
}

type baseMemoryOrderingInstructions interface {
	FENCE() // FENCE for data
}

type baseSystemInstructions interface {
	EBREAK() // Environment BREAK
	ECALL()  // Environment CALL
	MRET()   // Machine mode RETurn
	WFI()    // Wait For Interrupt
}
