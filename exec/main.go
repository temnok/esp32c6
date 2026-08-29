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
	var output []byte

	debug.Session(func(conn *debug.Conn) {
		conn.HartResetAndHalt(0)

		conn.WriteMem(0x4084_0000, code)

		conn.WriteGPR(isa.TP, 0x4080_0000)
		conn.WriteGPR(isa.SP, 0x4084_0000)
		conn.WriteCSR(csr.Mtvec, 0x4084_0000)
		conn.WriteCSR(csr.Dpc, 0x4084_0080)
		conn.WriteCSR(csr.Dcsr, 1<<csr.DcsrEbreakm|1<<csr.DcsrEbreaku|3<<csr.DcsrPrv)

		conn.HartResumeAndWaitForHalt(0)

		//fmt.Printf("dpc: %08X, mepc: %08X, mcause: %08X, mtvec: %08X\n",
		//	conn.ReadCSR(csr.Dpc), conn.ReadCSR(csr.Mepc), conn.ReadCSR(csr.Mcause), conn.ReadCSR(csr.Mtvec))
		//fmt.Printf("tp: %08X, sp: %08X\n", conn.ReadGPR(isa.TP), conn.ReadGPR(isa.SP))

		if tp := conn.ReadGPR(isa.TP); tp != 0 {
			outputLen := tp - 0x4080_0000
			output = make([]byte, outputLen)
			conn.ReadMem(0x4080_0000, output)
		}
	})

	fmt.Print(string(output))
}
