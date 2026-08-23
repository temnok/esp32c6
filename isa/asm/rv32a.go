package asm

func (asm *asm) AMOADD_W(rd, rs2, rs1 int)  { asm.R(0x0000202F, rd, rs1, rs2) }
func (asm *asm) AMOAND_W(rd, rs2, rs1 int)  { asm.R(0x6000202F, rd, rs1, rs2) }
func (asm *asm) AMOMAXU_W(rd, rs2, rs1 int) { asm.R(0xE000202F, rd, rs1, rs2) }
func (asm *asm) AMOMAX_W(rd, rs2, rs1 int)  { asm.R(0xA000202F, rd, rs1, rs2) }
func (asm *asm) AMOMINU_W(rd, rs2, rs1 int) { asm.R(0xC000202F, rd, rs1, rs2) }
func (asm *asm) AMOMIN_W(rd, rs2, rs1 int)  { asm.R(0x8000202F, rd, rs1, rs2) }
func (asm *asm) AMOOR_W(rd, rs2, rs1 int)   { asm.R(0x4000202F, rd, rs1, rs2) }
func (asm *asm) AMOSWAP_W(rd, rs2, rs1 int) { asm.R(0x0800202F, rd, rs1, rs2) }
func (asm *asm) AMOXOR_W(rd, rs2, rs1 int)  { asm.R(0x2000202F, rd, rs1, rs2) }
func (asm *asm) LR_W(rd, rs1 int)           { asm.R(0x1000202F, rd, rs1, 0) }
func (asm *asm) SC_W(rd, rs2, rs1 int)      { asm.R(0x1800202F, rd, rs1, rs2) }
