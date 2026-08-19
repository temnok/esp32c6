package isa

import (
	"github.com/deadsy/rvda"
	"testing"
)

func TestOpcodes(t *testing.T) {
	tests := map[uint]string{
		ADD:  "add zero,zero,zero",
		AND:  "and zero,zero,zero",
		OR:   "or zero,zero,zero",
		SLL:  "sll zero,zero,zero",
		SLT:  "slt zero,zero,zero",
		SLTU: "sltu zero,zero,zero",
		SRA:  "sra zero,zero,zero",
		SRL:  "srl zero,zero,zero",
		SUB:  "neg zero,zero",
		XOR:  "xor zero,zero,zero",

		ADDI:  "nop",
		ANDI:  "andi zero,zero,0",
		AUIPC: "auipc zero,0x0",
		LUI:   "lui zero,0x0",
		ORI:   "ori zero,zero,0",
		SLLI:  "slli zero,zero,0x0",
		SLTI:  "slti zero,zero,0",
		SLTIU: "sltiu zero,zero,0",
		SRAI:  "srai zero,zero,0x0",
		SRLI:  "srli zero,zero,0x0",
		XORI:  "xori zero,zero,0",

		BEQ:  "beqz zero,0",
		BGE:  "bgez zero,0",
		BGEU: "bgeu zero,zero,0",
		BLT:  "bltz zero,0",
		BLTU: "bltu zero,zero,0",
		BNE:  "bnez zero,0",
		JAL:  "j 0",
		JALR: "jalr zero,zero",

		LB:  "lb zero,0(zero)",
		LBU: "lbu zero,0(zero)",
		LH:  "lh zero,0(zero)",
		LHU: "lhu zero,0(zero)",
		LW:  "lw zero,0(zero)",
		SB:  "sb zero,0(zero)",
		SH:  "sh zero,0(zero)",
		SW:  "sw zero,0(zero)",

		FENCE: "fence",

		EBREAK: "ebreak",
		ECALL:  "ecall",

		MRET:       "mret",
		SFENCE_VMA: "sfence.vma",
		SRET:       "sret",
		URET:       "uret",
		WFI:        "wfi",
	}

	da, _ := rvda.New(32, rvda.ExtI|rvda.ExtM|rvda.ExtA|rvda.ExtC)

	for code, want := range tests {
		got := da.Disassemble(0, code).Assembly
		if got != want {
			t.Fatalf("Disassemble(%v):\n got %v\nwant %v", code, got, want)
		}
	}
}
