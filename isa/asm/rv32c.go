package asm

func (gen *gen) C_ADD(rd, rs2 int) { gen.CR(0x9002, rd, rs2) }
func (gen *gen) C_AND(rd, rs2 int) { gen.CA(0x8C61, rd, rs2) }
func (gen *gen) C_MV(rd, rs2 int)  { gen.CR(0x8002, rd, rs2) }
func (gen *gen) C_OR(rd, rs2 int)  { gen.CA(0x8C41, rd, rs2) }
func (gen *gen) C_SUB(rd, rs2 int) { gen.CA(0x8C01, rd, rs2) }
func (gen *gen) C_XOR(rd, rs2 int) { gen.CA(0x8C21, rd, rs2) }

func (gen *gen) C_ADDI(rd, imm int)     { gen.CI(0x0001, rd, imm) }
func (gen *gen) C_ADDI16SP(imm int)     { gen.CI16(0x6101, imm) }
func (gen *gen) C_ADDI4SPN(rd, imm int) { gen.CIW(0x0000, rd, imm) }
func (gen *gen) C_ANDI(rd, imm int)     { gen.CB(0x8801, rd, imm) }
func (gen *gen) C_LI(rd, imm int)       { gen.CI(0x4001, rd, imm) }
func (gen *gen) C_LUI(rd, imm int)      { gen.CI(0x6001, rd, imm) }
func (gen *gen) C_SLLI(rd, imm int)     { gen.CIsl(0x0002, rd, imm) }
func (gen *gen) C_SRAI(rd, imm int)     { gen.CBsr(0x8401, rd, imm) }
func (gen *gen) C_SRLI(rd, imm int)     { gen.CBsr(0x8001, rd, imm) }

func (gen *gen) C_BEQZ(rd, imm int) { gen.CB2(0xC001, rd, imm) }
func (gen *gen) C_BNEZ(rd, imm int) { gen.CB2(0xE001, rd, imm) }
func (gen *gen) C_J(imm int)        { gen.CJ(0xA001, imm) }
func (gen *gen) C_JAL(imm int)      { gen.CJ(0x2001, imm) }
func (gen *gen) C_JALR(rs1 int)     { gen.CR(0x9002, rs1, 0) }
func (gen *gen) C_JR(rs1 int)       { gen.CR(0x8002, rs1, 0) }

func (gen *gen) C_LW(rd, rs1, imm int)  { gen.CL(0x4000, rd, rs1, imm) }
func (gen *gen) C_LWSP(rd, imm int)     { gen.CI4(0x4002, rd, imm) }
func (gen *gen) C_SW(rs2, rs1, imm int) { gen.CS(0xC000, rs1, rs2, imm) }
func (gen *gen) C_SWSP(rs2, imm int)    { gen.CSS(0xC002, rs2, imm) }

func (gen *gen) C_EBREAK() { gen.CR(0x9002, 0, 0) }
func (gen *gen) C_NOP()    { gen.CI(0x0001, 0, 0) }
