package asm

func (asm *asm) URET() { asm.I(0x00200073, 0, 0, 0) }
