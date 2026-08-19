package opcode

func (gen *gen) CSRRC(rd, csr, rs1 int)  { gen.I(0x00003073, rd, rs1, csr) }
func (gen *gen) CSRRCI(rd, csr, imm int) { gen.I(0x00007073, rd, imm, csr) }
func (gen *gen) CSRRS(rd, csr, rs1 int)  { gen.I(0x00002073, rd, rs1, csr) }
func (gen *gen) CSRRSI(rd, csr, imm int) { gen.I(0x00006073, rd, imm, csr) }
func (gen *gen) CSRRW(rd, csr, rs1 int)  { gen.I(0x00001073, rd, rs1, csr) }
func (gen *gen) CSRRWI(rd, csr, imm int) { gen.I(0x00005073, rd, imm, csr) }
