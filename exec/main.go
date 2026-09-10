package main

import (
	"fmt"
	"github.com/temnok/esp32c6/check"
	"github.com/temnok/esp32c6/debug"
	"github.com/temnok/esp32c6/isa"
	"github.com/temnok/esp32c6/isa/csr"
	"log"
	"os"
)

func main() {
	defer check.RecoverAndPrintStack(log.Fatal)

	if len(os.Args) < 2 {
		fmt.Println("Expecting binary file path as argument")
		return
	}

	code := check.Err1(os.ReadFile(os.Args[1]))

	debug.Session(func(conn *debug.Conn) {
		conn.HartResetAndHalt(0)

		const (
			ramAddr  = 0x4080_0000
			appAddr  = ramAddr + 0x4_0000
			appEntry = appAddr + 0x80
		)

		conn.WriteMem(appAddr, code)

		conn.WriteGPR(isa.TP, ramAddr)
		conn.WriteGPR(isa.SP, appAddr)
		conn.WriteCSR(csr.Mtvec, appAddr)
		conn.WriteCSR(csr.Dpc, appEntry)
		conn.WriteCSR(csr.Dcsr, 1<<csr.DcsrEbreakm|1<<csr.DcsrEbreaku|3<<csr.DcsrPrv)
		conn.WriteCSR(csr.Mpcer, 1)
		conn.WriteWord(0x60008000+0x0048, 0) // Disable MWDT0 reset in TIMG_WDTCONFIG0_REG

		conn.HartResumeAndWaitForHalt(0)

		if tp := conn.ReadGPR(isa.TP); tp != 0 {
			outputLen := tp - ramAddr
			output := make([]byte, outputLen)
			conn.ReadMem(ramAddr, output)

			fmt.Print(string(output))
		}

		resetCause := conn.ReadWord(0x600B0400+0x0010) & 0x1F // LP_CLKRST_RESET_CAUSE_REG

		fmt.Printf("dpc: 0x%X, mepc: 0x%X, mcause: 0x%X, sp: 0x%X, reset_cause: 0x%X\n",
			conn.ReadCSR(csr.Dpc), conn.ReadCSR(csr.Mepc), conn.ReadCSR(csr.Mcause), conn.ReadGPR(isa.SP),
			resetCause)
	})
}
