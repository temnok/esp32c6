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

		conn.HartResumeAndWaitForHalt(0)

		if tp := conn.ReadGPR(isa.TP); tp != 0 {
			outputLen := tp - ramAddr
			output := make([]byte, outputLen)
			conn.ReadMem(ramAddr, output)

			fmt.Print(string(output))
		}

		fmt.Printf("dpc: 0x%X, mepc: 0x%X, mcause: 0x%X, sp: 0x%X\n",
			conn.ReadCSR(csr.Dpc), conn.ReadCSR(csr.Mepc), conn.ReadCSR(csr.Mcause), conn.ReadGPR(isa.SP))
	})
}
