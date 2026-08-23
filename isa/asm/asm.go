package asm

import (
	"fmt"
	"github.com/temnok/esp32c6/isa"
)

type Asm struct {
	isa.RV32IMACNZicsrZifencei

	curAddr   int
	labelAddr map[string]int
	code      []byte
	retry     bool
}

func (asm *Asm) Align(nBytes int) {
	mod := asm.curAddr % nBytes

	if mod != 0 {
		asm.curAddr += nBytes - mod
	}
}

func (asm *Asm) Skip(nBytes int) {
	asm.curAddr += nBytes
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

func (asm *Asm) pad() {
	for len(asm.code) < asm.curAddr {
		asm.code = append(asm.code, 0)
	}
}

func (asm *Asm) instr(opcode int) {
	asm.Align(2)
	asm.pad()

	asm.code = append(asm.code, byte(opcode), byte(opcode>>8))

	if opcode&3 == 3 {
		asm.code = append(asm.code, byte(opcode>>16), byte(opcode>>24))
	}

	asm.curAddr = len(asm.code)
}
