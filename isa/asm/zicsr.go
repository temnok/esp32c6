package asm

func (asm *asm) CSRRC(rd, csr, rs1 int)  { asm.I(0x00003073, rd, rs1, csr) }
func (asm *asm) CSRRCI(rd, csr, imm int) { asm.I(0x00007073, rd, imm, csr) }
func (asm *asm) CSRRS(rd, csr, rs1 int)  { asm.I(0x00002073, rd, rs1, csr) }
func (asm *asm) CSRRSI(rd, csr, imm int) { asm.I(0x00006073, rd, imm, csr) }
func (asm *asm) CSRRW(rd, csr, rs1 int)  { asm.I(0x00001073, rd, rs1, csr) }
func (asm *asm) CSRRWI(rd, csr, imm int) { asm.I(0x00005073, rd, imm, csr) }
