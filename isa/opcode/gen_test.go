package opcode

import (
	"github.com/deadsy/rvda"
	"github.com/stretchr/testify/assert"
	"github.com/temnok/esp32c6/isa"
	"testing"
)

var tests = map[string]func(isa.RV32){
	"add a0,a1,a2":  func(op isa.RV32) { op.ADD(10, 11, 12) },
	"and a0,a1,a2":  func(op isa.RV32) { op.AND(10, 11, 12) },
	"or a0,a1,a2":   func(op isa.RV32) { op.OR(10, 11, 12) },
	"sll a0,a1,a2":  func(op isa.RV32) { op.SLL(10, 11, 12) },
	"slt a0,a1,a2":  func(op isa.RV32) { op.SLT(10, 11, 12) },
	"sltu a0,a1,a2": func(op isa.RV32) { op.SLTU(10, 11, 12) },
	"sra a0,a1,a2":  func(op isa.RV32) { op.SRA(10, 11, 12) },
	"srl a0,a1,a2":  func(op isa.RV32) { op.SRL(10, 11, 12) },
	"sub a0,a1,a2":  func(op isa.RV32) { op.SUB(10, 11, 12) },
	"xor a0,a1,a2":  func(op isa.RV32) { op.XOR(10, 11, 12) },

	"addi a0,a1,12":  func(op isa.RV32) { op.ADDI(10, 11, 12) },
	"andi a0,a1,12":  func(op isa.RV32) { op.ANDI(10, 11, 12) },
	"auipc a1,0xc":   func(op isa.RV32) { op.AUIPC(11, 12) },
	"lui a1,0xc":     func(op isa.RV32) { op.LUI(11, 12) },
	"ori a0,a1,12":   func(op isa.RV32) { op.ORI(10, 11, 12) },
	"slli a0,a1,0xc": func(op isa.RV32) { op.SLLI(10, 11, 12) },
	"slti a0,a1,12":  func(op isa.RV32) { op.SLTI(10, 11, 12) },
	"sltiu a0,a1,12": func(op isa.RV32) { op.SLTIU(10, 11, 12) },
	"srai a0,a1,0xc": func(op isa.RV32) { op.SRAI(10, 11, 12) },
	"srli a0,a1,0xc": func(op isa.RV32) { op.SRLI(10, 11, 12) },
	"xori a0,a1,12":  func(op isa.RV32) { op.XORI(10, 11, 12) },

	"beq a0,a1,c":    func(op isa.RV32) { op.BEQ(10, 11, 12) },
	"bge a0,a1,c":    func(op isa.RV32) { op.BGE(10, 11, 12) },
	"bgeu a0,a1,c":   func(op isa.RV32) { op.BGEU(10, 11, 12) },
	"blt a0,a1,c":    func(op isa.RV32) { op.BLT(10, 11, 12) },
	"bltu a0,a1,c":   func(op isa.RV32) { op.BLTU(10, 11, 12) },
	"bne a0,a1,c":    func(op isa.RV32) { op.BNE(10, 11, 12) },
	"jal a1,c":       func(op isa.RV32) { op.JAL(11, 12) },
	"jalr a0,12(a1)": func(op isa.RV32) { op.JALR(10, 11, 12) },

	"lb a0,12(a1)":  func(op isa.RV32) { op.LB(10, 11, 12) },
	"lbu a0,12(a1)": func(op isa.RV32) { op.LBU(10, 11, 12) },
	"lh a0,12(a1)":  func(op isa.RV32) { op.LH(10, 11, 12) },
	"lhu a0,12(a1)": func(op isa.RV32) { op.LHU(10, 11, 12) },
	"lw a0,12(a1)":  func(op isa.RV32) { op.LW(10, 11, 12) },
	"sb a0,12(a1)":  func(op isa.RV32) { op.SB(10, 11, 12) },
	"sh a0,12(a1)":  func(op isa.RV32) { op.SH(10, 11, 12) },
	"sw a0,12(a1)":  func(op isa.RV32) { op.SW(10, 11, 12) },

	"fence": func(op isa.RV32) { op.FENCE() },

	"ebreak": func(op isa.RV32) { op.EBREAK() },
	"ecall":  func(op isa.RV32) { op.ECALL() },

	"div a0,a1,a2":    func(op isa.RV32) { op.DIV(10, 11, 12) },
	"divu a0,a1,a2":   func(op isa.RV32) { op.DIVU(10, 11, 12) },
	"mul a0,a1,a2":    func(op isa.RV32) { op.MUL(10, 11, 12) },
	"mulh a0,a1,a2":   func(op isa.RV32) { op.MULH(10, 11, 12) },
	"mulhsu a0,a1,a2": func(op isa.RV32) { op.MULHSU(10, 11, 12) },
	"mulhu a0,a1,a2":  func(op isa.RV32) { op.MULHU(10, 11, 12) },
	"rem a0,a1,a2":    func(op isa.RV32) { op.REM(10, 11, 12) },
	"remu a0,a1,a2":   func(op isa.RV32) { op.REMU(10, 11, 12) },

	"amoadd.w a0,a1,(a2)":  func(op isa.RV32) { op.AMOADD_W(10, 11, 12) },
	"amoand.w a0,a1,(a2)":  func(op isa.RV32) { op.AMOAND_W(10, 11, 12) },
	"amomax.w a0,a1,(a2)":  func(op isa.RV32) { op.AMOMAX_W(10, 11, 12) },
	"amomaxu.w a0,a1,(a2)": func(op isa.RV32) { op.AMOMAXU_W(10, 11, 12) },
	"amomin.w a0,a1,(a2)":  func(op isa.RV32) { op.AMOMIN_W(10, 11, 12) },
	"amominu.w a0,a1,(a2)": func(op isa.RV32) { op.AMOMINU_W(10, 11, 12) },
	"amoor.w a0,a1,(a2)":   func(op isa.RV32) { op.AMOOR_W(10, 11, 12) },
	"amoswap.w a0,a1,(a2)": func(op isa.RV32) { op.AMOSWAP_W(10, 11, 12) },
	"amoxor.w a0,a1,(a2)":  func(op isa.RV32) { op.AMOXOR_W(10, 11, 12) },
	"lr.w a1,(a2)":         func(op isa.RV32) { op.LR_W(11, 12) },
	"sc.w a0,a1,(a2)":      func(op isa.RV32) { op.SC_W(10, 11, 12) },

	"add a1,a1,a2": func(op isa.RV32) { op.C_ADD(11, 12) },
	"and a1,a1,a2": func(op isa.RV32) { op.C_AND(11, 12) },
	"mv a1,a2":     func(op isa.RV32) { op.C_MV(11, 12) },
	"or a1,a1,a2":  func(op isa.RV32) { op.C_OR(11, 12) },
	"sub a1,a1,a2": func(op isa.RV32) { op.C_SUB(11, 12) },
	"xor a1,a1,a2": func(op isa.RV32) { op.C_XOR(11, 12) },

	"addi a1,a1,12":   func(op isa.RV32) { op.C_ADDI(11, 12) },
	"addi sp,sp,16":   func(op isa.RV32) { op.C_ADDI16SP(16) },
	"addi sp,sp,32":   func(op isa.RV32) { op.C_ADDI16SP(32) },
	"addi sp,sp,64":   func(op isa.RV32) { op.C_ADDI16SP(64) },
	"addi sp,sp,128":  func(op isa.RV32) { op.C_ADDI16SP(128) },
	"addi sp,sp,256":  func(op isa.RV32) { op.C_ADDI16SP(256) },
	"addi sp,sp,-512": func(op isa.RV32) { op.C_ADDI16SP(512) },
	"addi a1,sp,4":    func(op isa.RV32) { op.C_ADDI4SPN(11, 4) },
	"addi a1,sp,8":    func(op isa.RV32) { op.C_ADDI4SPN(11, 8) },
	"addi a1,sp,16":   func(op isa.RV32) { op.C_ADDI4SPN(11, 16) },
	"addi a1,sp,32":   func(op isa.RV32) { op.C_ADDI4SPN(11, 32) },
	"addi a1,sp,64":   func(op isa.RV32) { op.C_ADDI4SPN(11, 64) },
	"addi a1,sp,128":  func(op isa.RV32) { op.C_ADDI4SPN(11, 128) },
	"addi a1,sp,256":  func(op isa.RV32) { op.C_ADDI4SPN(11, 256) },
	"addi a1,sp,512":  func(op isa.RV32) { op.C_ADDI4SPN(11, 512) },
	"andi a1,a1,12":   func(op isa.RV32) { op.C_ANDI(11, 12) },
	"li a1,12":        func(op isa.RV32) { op.C_LI(11, 12) },
	"lui a0,0xb":      func(op isa.RV32) { op.C_LUI(10, 11) },
	"slli a1,a1,0xc":  func(op isa.RV32) { op.C_SLLI(11, 12) },
	"srai a1,a1,0xc":  func(op isa.RV32) { op.C_SRAI(11, 12) },
	"srli a1,a1,0xc":  func(op isa.RV32) { op.C_SRLI(11, 12) },

	"beqz a1,2":    func(op isa.RV32) { op.C_BEQZ(11, 2) },
	"beqz a1,4":    func(op isa.RV32) { op.C_BEQZ(11, 4) },
	"beqz a1,8":    func(op isa.RV32) { op.C_BEQZ(11, 8) },
	"beqz a1,10":   func(op isa.RV32) { op.C_BEQZ(11, 16) },
	"beqz a1,20":   func(op isa.RV32) { op.C_BEQZ(11, 32) },
	"beqz a1,40":   func(op isa.RV32) { op.C_BEQZ(11, 64) },
	"beqz a1,80":   func(op isa.RV32) { op.C_BEQZ(11, 128) },
	"beqz a1,-100": func(op isa.RV32) { op.C_BEQZ(11, 256) },
	"bnez a1,2":    func(op isa.RV32) { op.C_BNEZ(11, 2) },
	"j 2":          func(op isa.RV32) { op.C_J(2) },
	"j 4":          func(op isa.RV32) { op.C_J(4) },
	"j 8":          func(op isa.RV32) { op.C_J(8) },
	"j 10":         func(op isa.RV32) { op.C_J(16) },
	"j 20":         func(op isa.RV32) { op.C_J(32) },
	"j 40":         func(op isa.RV32) { op.C_J(64) },
	"j 80":         func(op isa.RV32) { op.C_J(128) },
	"j 100":        func(op isa.RV32) { op.C_J(256) },
	"j 200":        func(op isa.RV32) { op.C_J(512) },
	"j 400":        func(op isa.RV32) { op.C_J(1024) },
	"j -800":       func(op isa.RV32) { op.C_J(2048) },
	"jal ra,2":     func(op isa.RV32) { op.C_JAL(2) },
	"jalr a0":      func(op isa.RV32) { op.C_JALR(10) },
	"jr a0":        func(op isa.RV32) { op.C_JR(10) },

	"lw a0,4(a1)":   func(op isa.RV32) { op.C_LW(10, 11, 4) },
	"lw a0,8(a1)":   func(op isa.RV32) { op.C_LW(10, 11, 8) },
	"lw a0,16(a1)":  func(op isa.RV32) { op.C_LW(10, 11, 16) },
	"lw a0,32(a1)":  func(op isa.RV32) { op.C_LW(10, 11, 32) },
	"lw a0,64(a1)":  func(op isa.RV32) { op.C_LW(10, 11, 64) },
	"lw a0,0(a1)":   func(op isa.RV32) { op.C_LW(10, 11, 128) },
	"lw a1,4(sp)":   func(op isa.RV32) { op.C_LWSP(11, 4) },
	"lw a1,8(sp)":   func(op isa.RV32) { op.C_LWSP(11, 8) },
	"lw a1,16(sp)":  func(op isa.RV32) { op.C_LWSP(11, 16) },
	"lw a1,32(sp)":  func(op isa.RV32) { op.C_LWSP(11, 32) },
	"lw a1,64(sp)":  func(op isa.RV32) { op.C_LWSP(11, 64) },
	"lw a1,128(sp)": func(op isa.RV32) { op.C_LWSP(11, 128) },
	"lw a1,0(sp)":   func(op isa.RV32) { op.C_LWSP(11, 256) },
	"sw a0,4(a1)":   func(op isa.RV32) { op.C_SW(10, 11, 4) },
	"sw a1,4(sp)":   func(op isa.RV32) { op.C_SWSP(11, 4) },
	"sw a1,8(sp)":   func(op isa.RV32) { op.C_SWSP(11, 8) },
	"sw a1,16(sp)":  func(op isa.RV32) { op.C_SWSP(11, 16) },
	"sw a1,32(sp)":  func(op isa.RV32) { op.C_SWSP(11, 32) },
	"sw a1,64(sp)":  func(op isa.RV32) { op.C_SWSP(11, 64) },
	"sw a1,128(sp)": func(op isa.RV32) { op.C_SWSP(11, 128) },
	"sw a1,0(sp)":   func(op isa.RV32) { op.C_SWSP(11, 256) },

	"ebreak TODO": func(op isa.RV32) { op.C_EBREAK() },
	"nop":         func(op isa.RV32) { op.C_NOP() },

	"mret":             func(op isa.RV32) { op.MRET() },
	"sfence.vma a2,a1": func(op isa.RV32) { op.SFENCE_VMA(11, 12) },
	"sret":             func(op isa.RV32) { op.SRET() },
	"uret":             func(op isa.RV32) { op.URET() },
	"wfi":              func(op isa.RV32) { op.WFI() },

	"csrrc a0,0x00b,a2":  func(op isa.RV32) { op.CSRRC(10, 11, 12) },
	"csrrci a0,0x00b,12": func(op isa.RV32) { op.CSRRCI(10, 11, 12) },
	"csrrs a0,0x00b,a2":  func(op isa.RV32) { op.CSRRS(10, 11, 12) },
	"csrrsi a0,0x00b,12": func(op isa.RV32) { op.CSRRSI(10, 11, 12) },
	"csrrw a0,0x00b,a2":  func(op isa.RV32) { op.CSRRW(10, 11, 12) },
	"csrrwi a0,0x00b,12": func(op isa.RV32) { op.CSRRWI(10, 11, 12) },

	"fence.i": func(op isa.RV32) { op.FENCE_I() },
}

func TestGen(t *testing.T) {
	da, _ := rvda.New(32, rvda.ExtI|rvda.ExtM|rvda.ExtA|rvda.ExtC)

	for want, op := range tests {
		var got string

		op(Gen(func(opcode int) {
			got = da.Disassemble(0, uint(opcode)).Assembly
		}))

		assert.Equal(t, want, got)
	}
}
