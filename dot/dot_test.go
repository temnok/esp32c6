package dot

import (
	"fmt"
	"github.com/deadsy/rvda"
	"github.com/temnok/esp32c6/isa"
	"github.com/temnok/esp32c6/isa/asm"
	"reflect"
	"testing"
)

func TestBlock(t *testing.T) {
	tests := []struct {
		input func(*Dot)
		want  []string
	}{
		{
			input: func(dot *Dot) {
				asm := &asm.Pseudo{asm.Asm(dot.Instr)}

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
			input: func(dot *Dot) {
				asm := &asm.Pseudo{asm.Asm(dot.Instr)}

				asm.C_J(dot.Offset("next"))
				dot.Label("next")
				asm.C_EBREAK()
			},
			want: []string{
				"0000 j 2",
				"0002 ebreak TODO",
			},
		},

		{
			input: func(dot *Dot) {
				asm := &asm.Pseudo{asm.Asm(dot.Instr)}

				dot.Label("start")
				asm.LA(isa.GP, dot.Offset("start"))
			},
			want: []string{
				"0000 auipc gp,0x0",
			},
		},

		{
			input: func(dot *Dot) {
				asm := &asm.Pseudo{asm.Asm(dot.Instr)}

				asm.C_NOP()
				dot.Label("start")
				asm.C_NOP()
				asm.LA(isa.GP, dot.Offset("start"))
				asm.LI(isa.A0, dot.Address("start"))
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

func disassemble(da *rvda.ISA, code []byte) []string {
	var disasm []string

	for i := 0; i+2 <= len(code); i += 2 {
		addr := uint(i)

		op := int(code[i+1])<<8 | int(code[i])
		if op&3 == 3 {
			op |= int(code[i+3])<<24 | int(code[i+2])<<16
			i += 2
		}

		instr := da.Disassemble(addr, uint(op))
		disasm = append(disasm, fmt.Sprintf("%04X %v", addr, instr.Assembly))
	}

	return disasm
}
