package opcode

import "github.com/temnok/esp32c6/isa"

func (asm *Asm) BEQZ(r, imm int)      { asm.BEQ(r, isa.Zero, imm) }
func (asm *Asm) BGEZ(r, imm int)      { asm.BGE(r, isa.Zero, imm) }
func (asm *Asm) BGT(ra, rb, imm int)  { asm.BLT(rb, ra, imm) }
func (asm *Asm) BGTU(ra, rb, imm int) { asm.BLTU(rb, ra, imm) }
func (asm *Asm) BGTZ(r, imm int)      { asm.BLT(isa.Zero, r, imm) }
func (asm *Asm) BLE(ra, rb, imm int)  { asm.BGE(rb, ra, imm) }
func (asm *Asm) BLEU(ra, rb, imm int) { asm.BGEU(rb, ra, imm) }
func (asm *Asm) BLEZ(r, imm int)      { asm.BGE(isa.Zero, r, imm) }
func (asm *Asm) BLTZ(r, imm int)      { asm.BLT(r, isa.Zero, imm) }
func (asm *Asm) BNEZ(r, imm int)      { asm.BNE(r, isa.Zero, imm) }
func (asm *Asm) J(imm int)            { asm.JAL(isa.Zero, imm) }
func (asm *Asm) JR(r int)             { asm.JALR(isa.Zero, r, 0) }
func (asm *Asm) NEG(rd, rs int)       { asm.SUB(rd, isa.Zero, rs) }
func (asm *Asm) NOT(rd, rs int)       { asm.XORI(rd, rs, -1) }
func (asm *Asm) RET()                 { asm.JALR(isa.Zero, isa.RA, 0) }
func (asm *Asm) SEQZ(rd, rs int)      { asm.SLTIU(rd, rs, 1) }
func (asm *Asm) SGTZ(rd, rs int)      { asm.SLT(rd, isa.Zero, rs) }
func (asm *Asm) SLTZ(rd, rs int)      { asm.SLT(rd, rs, isa.Zero) }
func (asm *Asm) SNEZ(rd, rs int)      { asm.SLTU(rd, isa.Zero, rs) }
