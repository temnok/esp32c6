package opcode

import "github.com/temnok/esp32c6/isa"

type Asm struct {
	isa.RV32IMACNZicsrZifencei

	curAddr   int
	labelAddr map[string]int
	code      []uint32
	retry     bool
}

func (asm *Asm) Label(name string) {
	if addr, known := asm.labelAddr[name]; known && addr != asm.curAddr {
		asm.retry = true
	}

	asm.labelAddr[name] = asm.curAddr
}

func (asm *Asm) Address(name string) int {
	return asm.addr(name, 0)
}

func (asm *Asm) Offset(name string) int {
	return asm.addr(name, asm.curAddr)
}

func (asm *Asm) addr(name string, curAddr int) int {
	if addr, known := asm.labelAddr[name]; known {
		return addr - curAddr
	}

	asm.retry = true
	return 0
}

func (asm *Asm) addInstr(opcode int) {
	op := uint32(opcode)
	compressed := op&3 != 3

	if asm.curAddr&3 == 0 {
		asm.code = append(asm.code, op)
	} else {
		asm.code[len(asm.code)-1] |= op << 16

		if !compressed {
			asm.code = append(asm.code, op>>16)
		}
	}

	if compressed {
		asm.curAddr += 2
	} else {
		asm.curAddr += 4
	}
}

func Assemble(block func(asm *Asm)) []uint32 {
	asm := &Asm{
		labelAddr: map[string]int{},
	}

	asm.RV32IMACNZicsrZifencei = Gen(asm.addInstr)

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
