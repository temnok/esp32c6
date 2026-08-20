package opcode

import "github.com/temnok/esp32c6/isa"

func Gen(callback func(int)) isa.RV32 {
	return &gen{callback}
}

type gen struct {
	callback func(int)
}

func (gen *gen) R(op, rd, rs1, rs2 int) {
	gen.callback(rs2&31<<20 | rs1&31<<15 | rd&31<<7 | op)
}

func (gen *gen) I(op, rd, rs1, imm int) {
	gen.callback(imm&0xFFF<<20 | rs1&31<<15 | rd&31<<7 | op)
}

func (gen *gen) S(op, rs2, rs1, imm int) {
	imm = imm>>5&0x7F<<25 | imm&0x1F<<7
	gen.callback(imm | rs2&31<<20 | rs1&31<<15 | op)
}

func (gen *gen) U(op, rd, imm int) {
	gen.callback(imm&0xFFFFF<<12 | rd&31<<7 | op)
}

func (gen *gen) B(op, rs1, rs2, imm int) {
	imm = imm>>12&1<<31 | imm>>5&0x3F<<25 | imm>>1&0xF<<8 | imm>>11&1<<7
	gen.callback(imm | rs2&31<<20 | rs1&31<<15 | op)
}

func (gen *gen) J(op, rd, imm int) {
	imm = imm>>20&1<<31 | imm>>1&0x3FF<<21 | imm>>11&1<<20 | imm>>12&0xFF<<12
	gen.callback(imm | rd&31<<7 | op)
}

func (gen *gen) CR(op, rd, rs2 int) {
	gen.callback(rd&31<<7 | rs2&31<<2 | op)
}

func (gen *gen) CA(op, rd, rs2 int) {
	gen.callback(rd&7<<7 | rs2&7<<2 | op)
}

func (gen *gen) CI(op, rd, imm int) {
	imm = imm>>5&1<<12 | imm&31<<2
	gen.callback(imm | rd&31<<7 | op)
}

func (gen *gen) CI4(op, rd, imm int) {
	imm = imm>>5&1<<12 | imm>>2&7<<4 | imm>>6&3<<2
	gen.callback(imm | rd&31<<7 | op)
}

func (gen *gen) CI16(op, imm int) {
	imm = imm>>9&1<<12 | imm>>4&1<<6 | imm>>6&1<<5 | imm>>7&3<<3 | imm>>5&1<<2
	gen.callback(imm | op)
}

func (gen *gen) CIW(op, rd, imm int) {
	imm = imm>>4&3<<11 | imm>>6&0xF<<7 | imm>>2&1<<6 | imm>>3&1<<5
	gen.callback(imm | rd&7<<2 | op)
}

func (gen *gen) CB(op, rd, imm int) {
	imm = imm>>5&1<<12 | imm&31<<2
	gen.callback(imm | rd&7<<7 | op)
}

func (gen *gen) CB2(op, rs1, imm int) {
	imm = imm>>8&1<<12 | imm>>3&3<<10 | imm>>6&3<<5 | imm>>1&3<<3 | imm>>5&1<<2
	gen.callback(imm | rs1&7<<7 | op)
}

func (gen *gen) CJ(op, imm int) {
	imm = imm>>11&1<<12 | imm>>4&1<<11 | imm>>8&3<<9 | imm>>10&1<<8 | imm>>6&1<<7 | imm>>7&1<<6 | imm>>1&7<<3 | imm>>5&1<<2
	gen.callback(imm | op)
}

func (gen *gen) CL(op, rd, rs1, imm int) {
	imm = imm>>3&7<<10 | imm>>2&1<<6 | imm>>6&1<<5
	gen.callback(imm | rs1&7<<7 | rd&7<<2 | op)
}

func (gen *gen) CS(op, rs1, rs2, imm int) {
	imm = imm>>3&7<<10 | imm>>2&1<<6 | imm>>6&1<<5
	gen.callback(imm | rs1&7<<7 | rs2&7<<2 | op)
}

func (gen *gen) CSS(op, rs2, imm int) {
	imm = imm>>2&15<<9 | imm>>6&3<<7
	gen.callback(imm | rs2&31<<2 | op)
}
