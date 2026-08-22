package asm

import (
	"fmt"
	"github.com/temnok/esp32c6/isa"
)

type Asm struct {
	isa.RV32IMACNZicsrZifencei

	curAddr   int
	labelAddr map[string]int
	code      []uint32
	retry     bool
}

func (asm *Asm) Label(label string) {
	if addr, known := asm.labelAddr[label]; known && addr != asm.curAddr {
		asm.retry = true
	}

	asm.labelAddr[label] = asm.curAddr
}

func (asm *Asm) Address(label string) int {
	return asm.addr(label, 0)
}

func (asm *Asm) Offset(label string) int {
	return asm.addr(label, asm.curAddr)
}

func (asm *Asm) addr(label string, curAddr int) int {
	if addr, known := asm.labelAddr[label]; known {
		return addr - curAddr
	}

	if asm.retry {
		panic(fmt.Errorf("unknown label: %v", label))
	}

	asm.retry = true
	return 0
}

func (asm *Asm) instr(opcode int) {
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
