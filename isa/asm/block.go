package asm

func Block(block func(asm *Asm)) []uint32 {
	asm := &Asm{
		labelAddr: map[string]int{},
	}

	asm.RV32IMACNZicsrZifencei = Gen(asm.instr)

	for {
		asm.curAddr = 0
		asm.code = asm.code[:0]
		asm.retry = false

		block(asm)

		if !asm.retry {
			break
		}
	}

	return asm.code
}
