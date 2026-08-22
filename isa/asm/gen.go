package asm

import (
	"fmt"
	"github.com/temnok/esp32c6/isa"
)

func Gen(callback func(int)) isa.RV32IMACNZicsrZifencei {
	return &gen{callback}
}

type gen struct {
	callback func(int)
}

func (gen *gen) R(op, rd, rs1, rs2 int) {
	assertRange(rd, 0, 32)
	assertRange(rs1, 0, 32)
	assertRange(rs2, 0, 32)

	gen.callback(rs2&31<<20 | rs1&31<<15 | rd&31<<7 | op)
}

func (gen *gen) I(op, rd, rs1, imm int) {
	assertRange(rd, 0, 32)
	assertRange(rs1, 0, 32)
	assertRange(imm, -0x800, 0x800)

	gen.callback(imm&0xFFF<<20 | rs1&31<<15 | rd&31<<7 | op)
}

func (gen *gen) Ish(op, rd, rs1, imm int) {
	assertRange(imm, 0, 32)

	gen.I(op, rd, rs1, imm)
}

func (gen *gen) S(op, rs2, rs1, imm int) {
	assertRange(rs1, 0, 32)
	assertRange(rs2, 0, 32)
	assertRange(imm, -0x800, 0x800)

	imm = imm>>5&0x7F<<25 | imm&0x1F<<7
	gen.callback(imm | rs2&31<<20 | rs1&31<<15 | op)
}

func (gen *gen) U(op, rd, imm int) {
	assertRange(rd, 0, 32)
	assertRange(imm, -0x80000, 0x100000)

	gen.callback(imm&0xFFFFF<<12 | rd&31<<7 | op)
}

func (gen *gen) B(op, rs1, rs2, imm int) {
	assertRange(rs1, 0, 32)
	assertRange(rs2, 0, 32)
	assertRange(imm, -0x1000, 0x1000)

	imm = imm>>12&1<<31 | imm>>5&0x3F<<25 | imm>>1&0xF<<8 | imm>>11&1<<7
	gen.callback(imm | rs2&31<<20 | rs1&31<<15 | op)
}

func (gen *gen) J(op, rd, imm int) {
	assertRange(rd, 0, 32)
	assertRange(imm, -0x200000, 0x200000)

	imm = imm>>20&1<<31 | imm>>1&0x3FF<<21 | imm>>11&1<<20 | imm>>12&0xFF<<12
	gen.callback(imm | rd&31<<7 | op)
}

func (gen *gen) CR(op, rd, rs2 int) {
	assertRange(rd, 0, 32)
	assertRange(rs2, 0, 32)

	gen.callback(rd&31<<7 | rs2&31<<2 | op)
}

func (gen *gen) CA(op, rd, rs2 int) {
	assertRange(rd, 8, 16)
	assertRange(rs2, 8, 16)

	gen.callback(rd&7<<7 | rs2&7<<2 | op)
}

func (gen *gen) CI(op, rd, imm int) {
	assertRange(rd, 0, 32)
	assertRange(imm, -0x20, 0x20)

	imm = imm>>5&1<<12 | imm&31<<2
	gen.callback(imm | rd&31<<7 | op)
}

func (gen *gen) CIsl(op, rd, imm int) {
	assertRange(imm, 0, 32)

	gen.CI(op, rd, imm)
}

func (gen *gen) CI4(op, rd, imm int) {
	assertRange(rd, 0, 32)
	assertRange(imm, 0, 0x100)
	assertAlign(imm, 4)

	imm = imm>>5&1<<12 | imm>>2&7<<4 | imm>>6&3<<2
	gen.callback(imm | rd&31<<7 | op)
}

func (gen *gen) CI16(op, imm int) {
	assertRange(imm, -0x200, 0x200)
	assertAlign(imm, 16)

	imm = imm>>9&1<<12 | imm>>4&1<<6 | imm>>6&1<<5 | imm>>7&3<<3 | imm>>5&1<<2
	gen.callback(imm | op)
}

func (gen *gen) CIW(op, rd, imm int) {
	assertRange(imm, 4, 0x400)
	assertAlign(imm, 4)

	imm = imm>>4&3<<11 | imm>>6&0xF<<7 | imm>>2&1<<6 | imm>>3&1<<5
	gen.callback(imm | rd&7<<2 | op)
}

func (gen *gen) CB(op, rd, imm int) {
	assertRange(rd, 8, 16)
	assertRange(imm, -0x20, 0x20)

	imm = imm>>5&1<<12 | imm&31<<2
	gen.callback(imm | rd&7<<7 | op)
}

func (gen *gen) CBsr(op, rd, imm int) {
	assertRange(imm, 0, 32)

	gen.CB(op, rd, imm)
}

func (gen *gen) CB2(op, rs1, imm int) {
	assertRange(rs1, 8, 16)
	assertRange(imm, -0x100, 0x100)
	assertAlign(imm, 2)

	imm = imm>>8&1<<12 | imm>>3&3<<10 | imm>>6&3<<5 | imm>>1&3<<3 | imm>>5&1<<2
	gen.callback(imm | rs1&7<<7 | op)
}

func (gen *gen) CJ(op, imm int) {
	assertRange(imm, -0x800, 0x800)
	assertAlign(imm, 2)

	imm = imm>>11&1<<12 | imm>>4&1<<11 | imm>>8&3<<9 | imm>>10&1<<8 | imm>>6&1<<7 | imm>>7&1<<6 | imm>>1&7<<3 | imm>>5&1<<2
	gen.callback(imm | op)
}

func (gen *gen) CL(op, rd, rs1, imm int) {
	assertRange(rd, 8, 16)
	assertRange(rs1, 8, 16)
	assertRange(imm, 0, 0x80)
	assertAlign(imm, 4)

	imm = imm>>3&7<<10 | imm>>2&1<<6 | imm>>6&1<<5
	gen.callback(imm | rs1&7<<7 | rd&7<<2 | op)
}

func (gen *gen) CS(op, rs1, rs2, imm int) {
	assertRange(rs1, 8, 16)
	assertRange(rs2, 8, 16)
	assertRange(imm, 0, 0x80)
	assertAlign(imm, 4)

	imm = imm>>3&7<<10 | imm>>2&1<<6 | imm>>6&1<<5
	gen.callback(imm | rs1&7<<7 | rs2&7<<2 | op)
}

func (gen *gen) CSS(op, rs2, imm int) {
	assertRange(rs2, 0, 32)
	assertRange(imm, 0, 0x100)
	assertAlign(imm, 4)

	imm = imm>>2&15<<9 | imm>>6&3<<7
	gen.callback(imm | rs2&31<<2 | op)
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
