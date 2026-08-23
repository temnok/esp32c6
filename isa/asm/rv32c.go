package asm

func (asm *asm) C_ADD(rd, rs2 int) { asm.CR(0x9002, rd, rs2) }
func (asm *asm) C_AND(rd, rs2 int) { asm.CA(0x8C61, rd, rs2) }
func (asm *asm) C_MV(rd, rs2 int)  { asm.CR(0x8002, rd, rs2) }
func (asm *asm) C_OR(rd, rs2 int)  { asm.CA(0x8C41, rd, rs2) }
func (asm *asm) C_SUB(rd, rs2 int) { asm.CA(0x8C01, rd, rs2) }
func (asm *asm) C_XOR(rd, rs2 int) { asm.CA(0x8C21, rd, rs2) }

func (asm *asm) C_ADDI(rd, imm int)     { asm.CI(0x0001, rd, imm) }
func (asm *asm) C_ADDI16SP(imm int)     { asm.CI16(0x6101, imm) }
func (asm *asm) C_ADDI4SPN(rd, imm int) { asm.CIW(0x0000, rd, imm) }
func (asm *asm) C_ANDI(rd, imm int)     { asm.CB(0x8801, rd, imm) }
func (asm *asm) C_LI(rd, imm int)       { asm.CI(0x4001, rd, imm) }
func (asm *asm) C_LUI(rd, imm int)      { asm.CI(0x6001, rd, imm) }
func (asm *asm) C_SLLI(rd, imm int)     { asm.CIsl(0x0002, rd, imm) }
func (asm *asm) C_SRAI(rd, imm int)     { asm.CBsr(0x8401, rd, imm) }
func (asm *asm) C_SRLI(rd, imm int)     { asm.CBsr(0x8001, rd, imm) }

func (asm *asm) C_BEQZ(rd, imm int) { asm.CB2(0xC001, rd, imm) }
func (asm *asm) C_BNEZ(rd, imm int) { asm.CB2(0xE001, rd, imm) }
func (asm *asm) C_J(imm int)        { asm.CJ(0xA001, imm) }
func (asm *asm) C_JAL(imm int)      { asm.CJ(0x2001, imm) }
func (asm *asm) C_JALR(rs1 int)     { asm.CR(0x9002, rs1, 0) }
func (asm *asm) C_JR(rs1 int)       { asm.CR(0x8002, rs1, 0) }

func (asm *asm) C_LW(rd, rs1, imm int)  { asm.CL(0x4000, rd, rs1, imm) }
func (asm *asm) C_LWSP(rd, imm int)     { asm.CI4(0x4002, rd, imm) }
func (asm *asm) C_SW(rs2, rs1, imm int) { asm.CS(0xC000, rs1, rs2, imm) }
func (asm *asm) C_SWSP(rs2, imm int)    { asm.CSS(0xC002, rs2, imm) }

func (asm *asm) C_EBREAK() { asm.CR(0x9002, 0, 0) }
func (asm *asm) C_NOP()    { asm.CI(0x0001, 0, 0) }
