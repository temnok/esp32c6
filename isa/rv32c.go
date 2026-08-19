package isa

type RV32C interface { //                    27 =
	compressedComputationalInstructions    // 6 +
	compressedComputationalImmInstructions // 9 +
	compressedControlTransferInstructions  // 6 +
	compressedLoadStoreInstructions        // 4 +
	compressedSpecialInstructions          // 2
}

type compressedComputationalInstructions interface {
	C_ADD(rd, rs2 int) // ADD
	C_AND(rd, rs2 int) // AND
	C_MV(rd, rs2 int)  // MoVe
	C_OR(rd, rs2 int)  // OR
	C_SUB(rd, rs2 int) // SUBtract
	C_XOR(rd, rs2 int) // eXclusive OR
}

type compressedComputationalImmInstructions interface {
	C_ADDI(rd, imm int)     // ADD Immediate
	C_ADDI16SP(imm int)     // ADD Immediate (multiples of 16) to Stack Pointer
	C_ADDI4SPN(rd, imm int) // ADD Immediate (multiples of 4) to Stack Pointer, Non-destructive
	C_ANDI(rd, imm int)     // AND with Immediate
	C_LI(rd, imm int)       // Load Immediate
	C_LUI(rd, imm int)      // Load Upper Immediate
	C_SLLI(rd, imm int)     // Shift Left, Logical by Immediate
	C_SRAI(rd, imm int)     // Shift Right, Arithmetic by Immediate
	C_SRLI(rd, imm int)     // Shift Right, Logical by Immediate
}

type compressedControlTransferInstructions interface {
	C_BEQZ(rs, offset int) // Branch if EQual to Zero
	C_BNEZ(rs, offset int) // Branch if Not Equal to Zero
	C_J(offset int)        // Jump by offset
	C_JAL(offset int)      // Jump And Link by offset
	C_JALR(rs int)         // Jump And Link by Register
	C_JR(rs int)           // Jump by Register
}

type compressedLoadStoreInstructions interface {
	C_LW(rd, rs, offset int)   // Load Word
	C_LWSP(rd, offset int)     // Load Word relative to Stack Pointer
	C_SW(rs2, rs1, offset int) // Store Word
	C_SWSP(rs, offset int)     // Store Word relative to Stack Pointer
}

type compressedSpecialInstructions interface {
	C_EBREAK() // Environment Break
	C_NOP()    // No-OP
}
