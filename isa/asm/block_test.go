package asm

import (
	"fmt"
	"github.com/deadsy/rvda"
	"github.com/temnok/esp32c6/isa"
	"reflect"
	"testing"
)

func TestBlock(t *testing.T) {
	tests := []struct {
		input func(*Asm)
		want  []string
	}{
		{
			input: func(asm *Asm) {
				asm.AUIPC(isa.GP, 0)
				asm.C_MV(isa.A0, isa.GP)
				asm.C_EBREAK()
			},
			want: []string{
				"0000 auipc gp,0x0",
				"0004 mv a0,gp",
				"0006 ebreak TODO",
			},
		},

		{
			input: func(asm *Asm) {
				asm.C_J(asm.Offset("next"))

				asm.Label("next")
				asm.C_EBREAK()
			},
			want: []string{
				"0000 j 2",
				"0002 ebreak TODO",
			},
		},

		{
			input: func(asm *Asm) {
				asm.Label("start")
				asm.LA(isa.GP, asm.Offset("start"))
			},
			want: []string{
				"0000 auipc gp,0x0",
			},
		},

		{
			input: func(asm *Asm) {
				asm.C_NOP()
				asm.Label("start")

				asm.C_NOP()
				asm.LA(isa.GP, asm.Offset("start"))
				asm.LI(isa.A0, asm.Address("start"))
				asm.LI(isa.A1, 0x12345678)
			},
			want: []string{
				"0000 nop",
				"0002 nop",
				"0004 auipc gp,0x0",
				"0008 addi gp,gp,-2",
				"000C li a0,2",
				"0010 lui a1,0x12345",
				"0014 addi a1,a1,1656",
			},
		},
	}

	da, _ := rvda.New(32, rvda.ExtI|rvda.ExtM|rvda.ExtA|rvda.ExtC)

	for _, test := range tests {
		code := Block(test.input)
		got := disassemble(da, code)

		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("\nwant %#v\n got %#v", test.want, got)
		}
	}
}

func disassemble(da *rvda.ISA, code []uint32) []string {
	var disasm []string

	for addr := 0; addr < 4*len(code); addr += 2 {
		op := code[addr>>2]

		var compressed bool
		if addr&3 == 0 {
			if compressed = op&3 != 3; compressed {
				op &= 0xFFFF
			}
		} else {
			op >>= 16
			if compressed = op&3 != 3; !compressed {
				op |= code[addr>>2+1] & 0xFFFF << 16
			}
		}

		instr := da.Disassemble(uint(addr), uint(op))
		disasm = append(disasm, fmt.Sprintf("%04X %v", addr, instr.Assembly))

		if !compressed {
			addr += 2
		}
	}

	return disasm
}
