package tests

import (
	"github.com/temnok/esp32c6/check"
	"github.com/temnok/esp32c6/debug"
	"github.com/temnok/esp32c6/dot"
	"github.com/temnok/esp32c6/isa"
	"github.com/temnok/esp32c6/isa/asm"
	"github.com/temnok/esp32c6/isa/csr"
	"testing"
)

func TestPerformance(t *testing.T) {
	defer check.RecoverAndPrintStack(t.Fatal)

	debug.Session(func(conn *debug.Conn) {
		conn.HartResetAndHalt(0)

		d := &dot.Dot{StartAddr: 0x4080_0000}
		asm := asm.AsmWithPseudo(d.Instr)

		assertPerf := func(testName string, expectedCycles int, block func()) {
			asm.CSRR(isa.T0, csr.Mpccr)

			block()

			asm.CSRR(isa.T1, csr.Mpccr)
			asm.SUB(isa.A0, isa.T1, isa.T0)
			asm.LI(isa.T1, expectedCycles)
			asm.BEQ(isa.A0, isa.T1, d.Offset(testName+"_ok"))

			d.Label(testName + "_fail")
			asm.C_EBREAK()

			d.Label(testName + "_ok")
		}

		d.Compile(func() {
			asm.LA(isa.A0, d.Offset("vector_base"))
			asm.CSRW(csr.Mtvec, isa.A0)
			asm.CSRWI(csr.Mpcer, 1)

			assertPerf("nop", 1, func() {
				asm.C_NOP()
			})

			assertPerf("mul", 1+10, func() {
				for range 10 {
					asm.MUL(isa.A3, isa.A4, isa.A5)
				}
			})

			assertPerf("mulh", 1+10*2, func() {
				for range 10 {
					asm.MULH(isa.A0, isa.A1, isa.A2)
				}
			})

			assertPerf("div", 1+10*10, func() {
				for range 10 {
					asm.DIV(isa.A3, isa.A4, isa.A5)
				}
			})

			assertPerf("clearmem", 2+512, func() {
				asm.LI(isa.A0, 0x4080_1000)
				for i := range 512 {
					asm.SW(isa.Zero, isa.A0, i*4)
				}
			})

			assertPerf("count", 6+2_000, func() {
				asm.LI(isa.A0, 1_000)

				d.Label("count_start")
				asm.C_ADDI(isa.A0, -1)
				asm.BNEZ(isa.A0, d.Offset("count_start"))
			})

			asm.C_NOP()
			d.Label("normal_exit")
			asm.C_EBREAK()

			d.Align(256)
			d.Label("vector_base")
			asm.C_EBREAK()
		})

		conn.WriteMem(d.StartAddr, d.Code)

		conn.WriteCSR(csr.Dcsr, 1<<csr.DcsrEbreakm|1<<csr.DcsrEbreaku|3<<csr.DcsrPrv)
		conn.WriteCSR(csr.Dpc, d.StartAddr)

		conn.HartResumeAndWaitForHalt(0)

		exitLabel := d.LabelByAddr(conn.ReadCSR(csr.Dpc))
		if exitLabel != "normal_exit" {
			t.Fatalf("%v, A0=%v", exitLabel, conn.ReadGPR(isa.A0))
		}
	})
}
