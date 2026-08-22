package asm

func (gen *gen) AMOADD_W(rd, rs2, rs1 int)  { gen.R(0x0000202F, rd, rs1, rs2) }
func (gen *gen) AMOAND_W(rd, rs2, rs1 int)  { gen.R(0x6000202F, rd, rs1, rs2) }
func (gen *gen) AMOMAXU_W(rd, rs2, rs1 int) { gen.R(0xE000202F, rd, rs1, rs2) }
func (gen *gen) AMOMAX_W(rd, rs2, rs1 int)  { gen.R(0xA000202F, rd, rs1, rs2) }
func (gen *gen) AMOMINU_W(rd, rs2, rs1 int) { gen.R(0xC000202F, rd, rs1, rs2) }
func (gen *gen) AMOMIN_W(rd, rs2, rs1 int)  { gen.R(0x8000202F, rd, rs1, rs2) }
func (gen *gen) AMOOR_W(rd, rs2, rs1 int)   { gen.R(0x4000202F, rd, rs1, rs2) }
func (gen *gen) AMOSWAP_W(rd, rs2, rs1 int) { gen.R(0x0800202F, rd, rs1, rs2) }
func (gen *gen) AMOXOR_W(rd, rs2, rs1 int)  { gen.R(0x2000202F, rd, rs1, rs2) }
func (gen *gen) LR_W(rd, rs1 int)           { gen.R(0x1000202F, rd, rs1, 0) }
func (gen *gen) SC_W(rd, rs2, rs1 int)      { gen.R(0x1800202F, rd, rs1, rs2) }
