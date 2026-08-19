package opcode

func (gen *gen) FENCE_I() { gen.I(0x0000100F, 0, 0, 0) }
