package asm

import "github.com/temnok/esp32c6/isa"

type Pseudo struct {
	isa.RV32IMACNZicsrZifencei
}

func (asm *Pseudo) LA(r, imm int) {
	hi, lo := imm>>12, int(int32(imm)<<20>>20)
	if lo < 0 {
		hi = (hi + 1) & 0xFFFFF
	}

	asm.AUIPC(r, hi)
	if lo != 0 {
		asm.ADDI(r, r, lo)
	}
}

func (asm *Pseudo) LI(r, imm int) {
	hi, lo := imm>>12, int(int32(imm)<<20>>20)
	if lo < 0 {
		hi = (hi + 1) & 0xFFFFF
	}

	if hi != 0 {
		asm.LUI(r, hi)
		if lo != 0 {
			asm.ADDI(r, r, lo)
		}
	} else {
		asm.ADDI(r, isa.Zero, lo)
	}
}

func (asm *Pseudo) BEQZ(r, imm int)      { asm.BEQ(r, isa.Zero, imm) }
func (asm *Pseudo) BGEZ(r, imm int)      { asm.BGE(r, isa.Zero, imm) }
func (asm *Pseudo) BGT(ra, rb, imm int)  { asm.BLT(rb, ra, imm) }
func (asm *Pseudo) BGTU(ra, rb, imm int) { asm.BLTU(rb, ra, imm) }
func (asm *Pseudo) BGTZ(r, imm int)      { asm.BLT(isa.Zero, r, imm) }
func (asm *Pseudo) BLE(ra, rb, imm int)  { asm.BGE(rb, ra, imm) }
func (asm *Pseudo) BLEU(ra, rb, imm int) { asm.BGEU(rb, ra, imm) }
func (asm *Pseudo) BLEZ(r, imm int)      { asm.BGE(isa.Zero, r, imm) }
func (asm *Pseudo) BLTZ(r, imm int)      { asm.BLT(r, isa.Zero, imm) }
func (asm *Pseudo) BNEZ(r, imm int)      { asm.BNE(r, isa.Zero, imm) }
func (asm *Pseudo) J(imm int)            { asm.JAL(isa.Zero, imm) }
func (asm *Pseudo) JR(r int)             { asm.JALR(isa.Zero, r, 0) }
func (asm *Pseudo) NEG(rd, rs int)       { asm.SUB(rd, isa.Zero, rs) }
func (asm *Pseudo) NOT(rd, rs int)       { asm.XORI(rd, rs, -1) }
func (asm *Pseudo) RET()                 { asm.JALR(isa.Zero, isa.RA, 0) }
func (asm *Pseudo) SEQZ(rd, rs int)      { asm.SLTIU(rd, rs, 1) }
func (asm *Pseudo) SGTZ(rd, rs int)      { asm.SLT(rd, isa.Zero, rs) }
func (asm *Pseudo) SLTZ(rd, rs int)      { asm.SLT(rd, rs, isa.Zero) }
func (asm *Pseudo) SNEZ(rd, rs int)      { asm.SLTU(rd, isa.Zero, rs) }

func (asm *Pseudo) CSRC(csr, rs int)   { asm.CSRRC(isa.Zero, csr, rs) }
func (asm *Pseudo) CSRCI(csr, imm int) { asm.CSRRCI(isa.Zero, csr, imm) }
func (asm *Pseudo) CSRR(rd, csr int)   { asm.CSRRS(rd, csr, isa.Zero) }
func (asm *Pseudo) CSRS(csr, rs int)   { asm.CSRRS(isa.Zero, csr, rs) }
func (asm *Pseudo) CSRSI(csr, imm int) { asm.CSRRSI(isa.Zero, csr, imm) }
func (asm *Pseudo) CSRW(csr, rs int)   { asm.CSRRW(isa.Zero, csr, rs) }
func (asm *Pseudo) CSRWI(csr, imm int) { asm.CSRRWI(isa.Zero, csr, imm) }
