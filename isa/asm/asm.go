package asm

import (
	"fmt"
	"github.com/temnok/esp32c6/isa"
)

func Asm(instr func(int)) isa.RV32IMACNZicsrZifencei {
	return &asm{instr}
}

type asm struct {
	instr func(int)
}

func (asm *asm) R(op, rd, rs1, rs2 int) {
	assertRange(rd, 0, 32)
	assertRange(rs1, 0, 32)
	assertRange(rs2, 0, 32)

	asm.instr(rs2&31<<20 | rs1&31<<15 | rd&31<<7 | op)
}

func (asm *asm) I(op, rd, rs1, imm int) {
	assertRange(rd, 0, 32)
	assertRange(rs1, 0, 32)
	assertRange(imm, -0x800, 0x800)

	asm.instr(imm&0xFFF<<20 | rs1&31<<15 | rd&31<<7 | op)
}

func (asm *asm) Ish(op, rd, rs1, imm int) {
	assertRange(imm, 0, 32)

	asm.I(op, rd, rs1, imm)
}

func (asm *asm) S(op, rs2, rs1, imm int) {
	assertRange(rs1, 0, 32)
	assertRange(rs2, 0, 32)
	assertRange(imm, -0x800, 0x800)

	imm = imm>>5&0x7F<<25 | imm&0x1F<<7
	asm.instr(imm | rs2&31<<20 | rs1&31<<15 | op)
}

func (asm *asm) U(op, rd, imm int) {
	assertRange(rd, 0, 32)
	assertRange(imm, -0x80000, 0x100000)

	asm.instr(imm&0xFFFFF<<12 | rd&31<<7 | op)
}

func (asm *asm) B(op, rs1, rs2, imm int) {
	assertRange(rs1, 0, 32)
	assertRange(rs2, 0, 32)
	assertRange(imm, -0x1000, 0x1000)

	imm = imm>>12&1<<31 | imm>>5&0x3F<<25 | imm>>1&0xF<<8 | imm>>11&1<<7
	asm.instr(imm | rs2&31<<20 | rs1&31<<15 | op)
}

func (asm *asm) J(op, rd, imm int) {
	assertRange(rd, 0, 32)
	assertRange(imm, -0x200000, 0x200000)

	imm = imm>>20&1<<31 | imm>>1&0x3FF<<21 | imm>>11&1<<20 | imm>>12&0xFF<<12
	asm.instr(imm | rd&31<<7 | op)
}

func (asm *asm) CR(op, rd, rs2 int) {
	assertRange(rd, 0, 32)
	assertRange(rs2, 0, 32)

	asm.instr(rd&31<<7 | rs2&31<<2 | op)
}

func (asm *asm) CA(op, rd, rs2 int) {
	assertRange(rd, 8, 16)
	assertRange(rs2, 8, 16)

	asm.instr(rd&7<<7 | rs2&7<<2 | op)
}

func (asm *asm) CI(op, rd, imm int) {
	assertRange(rd, 0, 32)
	assertRange(imm, -0x20, 0x20)

	imm = imm>>5&1<<12 | imm&31<<2
	asm.instr(imm | rd&31<<7 | op)
}

func (asm *asm) CIsl(op, rd, imm int) {
	assertRange(imm, 0, 32)

	asm.CI(op, rd, imm)
}

func (asm *asm) CI4(op, rd, imm int) {
	assertRange(rd, 0, 32)
	assertRange(imm, 0, 0x100)
	assertAlign(imm, 4)

	imm = imm>>5&1<<12 | imm>>2&7<<4 | imm>>6&3<<2
	asm.instr(imm | rd&31<<7 | op)
}

func (asm *asm) CI16(op, imm int) {
	assertRange(imm, -0x200, 0x200)
	assertAlign(imm, 16)

	imm = imm>>9&1<<12 | imm>>4&1<<6 | imm>>6&1<<5 | imm>>7&3<<3 | imm>>5&1<<2
	asm.instr(imm | op)
}

func (asm *asm) CIW(op, rd, imm int) {
	assertRange(imm, 4, 0x400)
	assertAlign(imm, 4)

	imm = imm>>4&3<<11 | imm>>6&0xF<<7 | imm>>2&1<<6 | imm>>3&1<<5
	asm.instr(imm | rd&7<<2 | op)
}

func (asm *asm) CB(op, rd, imm int) {
	assertRange(rd, 8, 16)
	assertRange(imm, -0x20, 0x20)

	imm = imm>>5&1<<12 | imm&31<<2
	asm.instr(imm | rd&7<<7 | op)
}

func (asm *asm) CBsr(op, rd, imm int) {
	assertRange(imm, 0, 32)

	asm.CB(op, rd, imm)
}

func (asm *asm) CB2(op, rs1, imm int) {
	assertRange(rs1, 8, 16)
	assertRange(imm, -0x100, 0x100)
	assertAlign(imm, 2)

	imm = imm>>8&1<<12 | imm>>3&3<<10 | imm>>6&3<<5 | imm>>1&3<<3 | imm>>5&1<<2
	asm.instr(imm | rs1&7<<7 | op)
}

func (asm *asm) CJ(op, imm int) {
	assertRange(imm, -0x800, 0x800)
	assertAlign(imm, 2)

	imm = imm>>11&1<<12 | imm>>4&1<<11 | imm>>8&3<<9 | imm>>10&1<<8 | imm>>6&1<<7 | imm>>7&1<<6 | imm>>1&7<<3 | imm>>5&1<<2
	asm.instr(imm | op)
}

func (asm *asm) CL(op, rd, rs1, imm int) {
	assertRange(rd, 8, 16)
	assertRange(rs1, 8, 16)
	assertRange(imm, 0, 0x80)
	assertAlign(imm, 4)

	imm = imm>>3&7<<10 | imm>>2&1<<6 | imm>>6&1<<5
	asm.instr(imm | rs1&7<<7 | rd&7<<2 | op)
}

func (asm *asm) CS(op, rs1, rs2, imm int) {
	assertRange(rs1, 8, 16)
	assertRange(rs2, 8, 16)
	assertRange(imm, 0, 0x80)
	assertAlign(imm, 4)

	imm = imm>>3&7<<10 | imm>>2&1<<6 | imm>>6&1<<5
	asm.instr(imm | rs1&7<<7 | rs2&7<<2 | op)
}

func (asm *asm) CSS(op, rs2, imm int) {
	assertRange(rs2, 0, 32)
	assertRange(imm, 0, 0x100)
	assertAlign(imm, 4)

	imm = imm>>2&15<<9 | imm>>6&3<<7
	asm.instr(imm | rs2&31<<2 | op)
}

func assertRange(val, min, max int) {
	if val < min || max <= val {
		panic(fmt.Errorf("value %v must be in range [%v..%v]", val, min, max-1))
	}
}

func assertAlign(val, align int) {
	if val&(align-1) != 0 {
		panic(fmt.Errorf("value %v should be divisible by %v", val, align))
	}
}
