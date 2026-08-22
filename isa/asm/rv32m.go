package asm

func (gen *gen) DIV(rd, rs1, rs2 int)    { gen.R(0x02004033, rd, rs1, rs2) }
func (gen *gen) DIVU(rd, rs1, rs2 int)   { gen.R(0x02005033, rd, rs1, rs2) }
func (gen *gen) MUL(rd, rs1, rs2 int)    { gen.R(0x02000033, rd, rs1, rs2) }
func (gen *gen) MULH(rd, rs1, rs2 int)   { gen.R(0x02001033, rd, rs1, rs2) }
func (gen *gen) MULHSU(rd, rs1, rs2 int) { gen.R(0x02002033, rd, rs1, rs2) }
func (gen *gen) MULHU(rd, rs1, rs2 int)  { gen.R(0x02003033, rd, rs1, rs2) }
func (gen *gen) REM(rd, rs1, rs2 int)    { gen.R(0x02006033, rd, rs1, rs2) }
func (gen *gen) REMU(rd, rs1, rs2 int)   { gen.R(0x02007033, rd, rs1, rs2) }
