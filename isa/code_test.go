package isa

import (
	"github.com/deadsy/rvda"
	"testing"
)

func TestCodes(t *testing.T) {
	tests := map[uint]string{
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
