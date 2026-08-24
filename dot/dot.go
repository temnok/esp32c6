package dot

import (
	"fmt"
)

type Dot struct {
	StartAddr int
	LabelAddr map[string]int
	Code      []byte

	curAddr int
	retry   bool
}

func (dot *Dot) Compile(block func()) {
	dot.LabelAddr = map[string]int{}

	for dot.retry = true; dot.retry; {
		dot.retry = false
		dot.curAddr = dot.StartAddr
		dot.Code = dot.Code[:0]

		block()
	}
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
	if addr, known := dot.LabelAddr[label]; known && addr != dot.curAddr {
		dot.retry = true
	}

	dot.LabelAddr[label] = dot.curAddr
}

func (dot *Dot) Address(label string) int {
	if addr, known := dot.LabelAddr[label]; known {
		return addr
	}

	if dot.retry {
		panic(fmt.Errorf("unknown label: %v", label))
	}

	dot.retry = true
	return 0
}

func (dot *Dot) Offset(label string) int {
	return dot.Address(label) - dot.curAddr
}

func (dot *Dot) pad() {
	for end := dot.curAddr - dot.StartAddr; len(dot.Code) < end; {
		dot.Code = append(dot.Code, 0)
	}
}

func (dot *Dot) Instr(opcode int) {
	dot.Align(2)
	dot.pad()

	dot.Code = append(dot.Code, byte(opcode), byte(opcode>>8))

	if opcode&3 == 3 {
		dot.Code = append(dot.Code, byte(opcode>>16), byte(opcode>>24))
	}

	dot.curAddr = dot.StartAddr + len(dot.Code)
}
