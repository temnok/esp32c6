package dot

import (
	"fmt"
)

type Dot struct {
	curAddr   int
	labelAddr map[string]int
	code      []byte
	retry     bool
}

func (dot *Dot) Align(nBytes int) {
	mod := dot.curAddr % nBytes

	if mod != 0 {
		dot.curAddr += nBytes - mod
	}
}

func (dot *Dot) Skip(nBytes int) {
	dot.curAddr += nBytes
}

func (dot *Dot) Label(label string) {
	if addr, known := dot.labelAddr[label]; known && addr != dot.curAddr {
		dot.retry = true
	}

	dot.labelAddr[label] = dot.curAddr
}

func (dot *Dot) Address(label string) int {
	return dot.addr(label, 0)
}

func (dot *Dot) Offset(label string) int {
	return dot.addr(label, dot.curAddr)
}

func (dot *Dot) addr(label string, curAddr int) int {
	if addr, known := dot.labelAddr[label]; known {
		return addr - curAddr
	}

	if dot.retry {
		panic(fmt.Errorf("unknown label: %v", label))
	}

	dot.retry = true
	return 0
}

func (dot *Dot) pad() {
	for len(dot.code) < dot.curAddr {
		dot.code = append(dot.code, 0)
	}
}

func (dot *Dot) Instr(opcode int) {
	dot.Align(2)
	dot.pad()

	dot.code = append(dot.code, byte(opcode), byte(opcode>>8))

	if opcode&3 == 3 {
		dot.code = append(dot.code, byte(opcode>>16), byte(opcode>>24))
	}

	dot.curAddr = len(dot.code)
}

func (dot *Dot) Block(block func()) {
	dot.labelAddr = map[string]int{}

	for {
		dot.curAddr = 0
		dot.code = dot.code[:0]
		dot.retry = false

		block()

		if !dot.retry {
			break
		}
	}
}

func (dot *Dot) Code() []byte {
	return dot.code
}
