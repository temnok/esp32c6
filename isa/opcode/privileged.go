package opcode

func (gen *gen) MRET()                   { gen.I(0x30200073, 0, 0, 0) }
func (gen *gen) SFENCE_VMA(rs1, rs2 int) { gen.I(0x12000073, 0, rs1, rs2) }
func (gen *gen) SRET()                   { gen.I(0x10200073, 0, 0, 0) }
func (gen *gen) URET()                   { gen.I(0x00200073, 0, 0, 0) }
func (gen *gen) WFI()                    { gen.I(0x10500073, 0, 0, 0) }
