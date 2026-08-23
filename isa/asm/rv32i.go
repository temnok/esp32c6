package asm

func (asm *asm) ADD(rd, rs1, rs2 int)  { asm.R(0x00000033, rd, rs1, rs2) }
func (asm *asm) AND(rd, rs1, rs2 int)  { asm.R(0x00007033, rd, rs1, rs2) }
func (asm *asm) OR(rd, rs1, rs2 int)   { asm.R(0x00006033, rd, rs1, rs2) }
func (asm *asm) SLL(rd, rs1, rs2 int)  { asm.R(0x00001033, rd, rs1, rs2) }
func (asm *asm) SLT(rd, rs1, rs2 int)  { asm.R(0x00002033, rd, rs1, rs2) }
func (asm *asm) SLTU(rd, rs1, rs2 int) { asm.R(0x00003033, rd, rs1, rs2) }
func (asm *asm) SRA(rd, rs1, rs2 int)  { asm.R(0x40005033, rd, rs1, rs2) }
func (asm *asm) SRL(rd, rs1, rs2 int)  { asm.R(0x00005033, rd, rs1, rs2) }
func (asm *asm) SUB(rd, rs1, rs2 int)  { asm.R(0x40000033, rd, rs1, rs2) }
func (asm *asm) XOR(rd, rs1, rs2 int)  { asm.R(0x00004033, rd, rs1, rs2) }

func (asm *asm) ADDI(rd, rs1, imm int)  { asm.I(0x00000013, rd, rs1, imm) }
func (asm *asm) ANDI(rd, rs1, imm int)  { asm.I(0x00007013, rd, rs1, imm) }
func (asm *asm) AUIPC(rd, imm int)      { asm.U(0x00000017, rd, imm) }
func (asm *asm) LUI(rd, imm int)        { asm.U(0x00000037, rd, imm) }
func (asm *asm) ORI(rd, rs1, imm int)   { asm.I(0x00006013, rd, rs1, imm) }
func (asm *asm) SLLI(rd, rs1, imm int)  { asm.Ish(0x00001013, rd, rs1, imm) }
func (asm *asm) SLTI(rd, rs1, imm int)  { asm.I(0x00002013, rd, rs1, imm) }
func (asm *asm) SLTIU(rd, rs1, imm int) { asm.I(0x00003013, rd, rs1, imm) }
func (asm *asm) SRAI(rd, rs1, imm int)  { asm.Ish(0x40005013, rd, rs1, imm) }
func (asm *asm) SRLI(rd, rs1, imm int)  { asm.Ish(0x00005013, rd, rs1, imm) }
func (asm *asm) XORI(rd, rs1, imm int)  { asm.I(0x00004013, rd, rs1, imm) }

func (asm *asm) BEQ(rs1, rs2, imm int)  { asm.B(0x00000063, rs1, rs2, imm) }
func (asm *asm) BGE(rs1, rs2, imm int)  { asm.B(0x00005063, rs1, rs2, imm) }
func (asm *asm) BGEU(rs1, rs2, imm int) { asm.B(0x00007063, rs1, rs2, imm) }
func (asm *asm) BLT(rs1, rs2, imm int)  { asm.B(0x00004063, rs1, rs2, imm) }
func (asm *asm) BLTU(rs1, rs2, imm int) { asm.B(0x00006063, rs1, rs2, imm) }
func (asm *asm) BNE(rs1, rs2, imm int)  { asm.B(0x00001063, rs1, rs2, imm) }
func (asm *asm) JAL(rd, imm int)        { asm.J(0x0000006F, rd, imm) }
func (asm *asm) JALR(rd, rs1, imm int)  { asm.I(0x00000067, rd, rs1, imm) }

func (asm *asm) LB(rd, rs1, imm int)  { asm.I(0x00000003, rd, rs1, imm) }
func (asm *asm) LBU(rd, rs1, imm int) { asm.I(0x00004003, rd, rs1, imm) }
func (asm *asm) LH(rd, rs1, imm int)  { asm.I(0x00001003, rd, rs1, imm) }
func (asm *asm) LHU(rd, rs1, imm int) { asm.I(0x00005003, rd, rs1, imm) }
func (asm *asm) LW(rd, rs1, imm int)  { asm.I(0x00002003, rd, rs1, imm) }
func (asm *asm) SB(rs2, rs1, imm int) { asm.S(0x00000023, rs2, rs1, imm) }
func (asm *asm) SH(rs2, rs1, imm int) { asm.S(0x00001023, rs2, rs1, imm) }
func (asm *asm) SW(rs2, rs1, imm int) { asm.S(0x00002023, rs2, rs1, imm) }

func (asm *asm) FENCE() { asm.I(0x0000000F, 0, 0, 0) }

func (asm *asm) EBREAK() { asm.I(0x00100073, 0, 0, 0) }
func (asm *asm) ECALL()  { asm.I(0x00000073, 0, 0, 0) }
func (asm *asm) MRET()   { asm.I(0x30200073, 0, 0, 0) }
func (asm *asm) WFI()    { asm.I(0x10500073, 0, 0, 0) }
