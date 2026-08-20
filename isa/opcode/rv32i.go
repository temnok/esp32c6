package opcode

func (gen *gen) ADD(rd, rs1, rs2 int)  { gen.R(0x00000033, rd, rs1, rs2) }
func (gen *gen) AND(rd, rs1, rs2 int)  { gen.R(0x00007033, rd, rs1, rs2) }
func (gen *gen) OR(rd, rs1, rs2 int)   { gen.R(0x00006033, rd, rs1, rs2) }
func (gen *gen) SLL(rd, rs1, rs2 int)  { gen.R(0x00001033, rd, rs1, rs2) }
func (gen *gen) SLT(rd, rs1, rs2 int)  { gen.R(0x00002033, rd, rs1, rs2) }
func (gen *gen) SLTU(rd, rs1, rs2 int) { gen.R(0x00003033, rd, rs1, rs2) }
func (gen *gen) SRA(rd, rs1, rs2 int)  { gen.R(0x40005033, rd, rs1, rs2) }
func (gen *gen) SRL(rd, rs1, rs2 int)  { gen.R(0x00005033, rd, rs1, rs2) }
func (gen *gen) SUB(rd, rs1, rs2 int)  { gen.R(0x40000033, rd, rs1, rs2) }
func (gen *gen) XOR(rd, rs1, rs2 int)  { gen.R(0x00004033, rd, rs1, rs2) }

func (gen *gen) ADDI(rd, rs1, imm int)  { gen.I(0x00000013, rd, rs1, imm) }
func (gen *gen) ANDI(rd, rs1, imm int)  { gen.I(0x00007013, rd, rs1, imm) }
func (gen *gen) AUIPC(rd, imm int)      { gen.U(0x00000017, rd, imm) }
func (gen *gen) LUI(rd, imm int)        { gen.U(0x00000037, rd, imm) }
func (gen *gen) ORI(rd, rs1, imm int)   { gen.I(0x00006013, rd, rs1, imm) }
func (gen *gen) SLLI(rd, rs1, imm int)  { gen.I(0x00001013, rd, rs1, imm&31) }
func (gen *gen) SLTI(rd, rs1, imm int)  { gen.I(0x00002013, rd, rs1, imm) }
func (gen *gen) SLTIU(rd, rs1, imm int) { gen.I(0x00003013, rd, rs1, imm) }
func (gen *gen) SRAI(rd, rs1, imm int)  { gen.I(0x40005013, rd, rs1, imm&31) }
func (gen *gen) SRLI(rd, rs1, imm int)  { gen.I(0x00005013, rd, rs1, imm&31) }
func (gen *gen) XORI(rd, rs1, imm int)  { gen.I(0x00004013, rd, rs1, imm) }

func (gen *gen) BEQ(rs1, rs2, imm int)  { gen.B(0x00000063, rs1, rs2, imm) }
func (gen *gen) BGE(rs1, rs2, imm int)  { gen.B(0x00005063, rs1, rs2, imm) }
func (gen *gen) BGEU(rs1, rs2, imm int) { gen.B(0x00007063, rs1, rs2, imm) }
func (gen *gen) BLT(rs1, rs2, imm int)  { gen.B(0x00004063, rs1, rs2, imm) }
func (gen *gen) BLTU(rs1, rs2, imm int) { gen.B(0x00006063, rs1, rs2, imm) }
func (gen *gen) BNE(rs1, rs2, imm int)  { gen.B(0x00001063, rs1, rs2, imm) }
func (gen *gen) JAL(rd, imm int)        { gen.J(0x0000006F, rd, imm) }
func (gen *gen) JALR(rd, rs1, imm int)  { gen.I(0x00000067, rd, rs1, imm) }

func (gen *gen) LB(rd, rs1, imm int)  { gen.I(0x00000003, rd, rs1, imm) }
func (gen *gen) LBU(rd, rs1, imm int) { gen.I(0x00004003, rd, rs1, imm) }
func (gen *gen) LH(rd, rs1, imm int)  { gen.I(0x00001003, rd, rs1, imm) }
func (gen *gen) LHU(rd, rs1, imm int) { gen.I(0x00005003, rd, rs1, imm) }
func (gen *gen) LW(rd, rs1, imm int)  { gen.I(0x00002003, rd, rs1, imm) }
func (gen *gen) SB(rs2, rs1, imm int) { gen.S(0x00000023, rs2, rs1, imm) }
func (gen *gen) SH(rs2, rs1, imm int) { gen.S(0x00001023, rs2, rs1, imm) }
func (gen *gen) SW(rs2, rs1, imm int) { gen.S(0x00002023, rs2, rs1, imm) }

func (gen *gen) FENCE() { gen.I(0x0000000F, 0, 0, 0) }

func (gen *gen) EBREAK() { gen.I(0x00100073, 0, 0, 0) }
func (gen *gen) ECALL()  { gen.I(0x00000073, 0, 0, 0) }
