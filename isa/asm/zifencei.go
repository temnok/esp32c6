package asm

func (asm *asm) FENCE_I() { asm.I(0x0000100F, 0, 0, 0) }
