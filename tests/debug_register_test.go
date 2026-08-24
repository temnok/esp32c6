package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/temnok/esp32c6/check"
	"github.com/temnok/esp32c6/debug"
	"github.com/temnok/esp32c6/isa"
	"github.com/temnok/esp32c6/isa/csr"
	"testing"
)

func TestRegistersAfterReset(t *testing.T) {
	defer check.RecoverAndPrintStack(t.Fatal)

	debug.Session(func(conn *debug.Conn) {
		conn.HartResetAndHalt(0)

		assert.Equal(t, 0, conn.ReadCSR(csr.Ustatus))
		assert.Equal(t, 0, conn.ReadCSR(csr.Uie))
		assert.Equal(t, 1, conn.ReadCSR(csr.Utvec))
		//assert.Equal(t, 0, conn.ReadCSR(csr.Uscratch)) // ???
		assert.Equal(t, 0, conn.ReadCSR(csr.Uepc))
		assert.Equal(t, 0, conn.ReadCSR(csr.Ucause))
		assert.Equal(t, 0, conn.ReadCSR(csr.Uip))

		assert.Equal(t, 1<<csr.MstatusTW|3<<csr.MstatusMPP, conn.ReadCSR(csr.Mstatus))
		misa := 1<<30 | 1<<('n'-'a') | 1<<('u'-'a') | 1<<('x'-'a') |
			1<<('i'-'a') | 1<<('m'-'a') | 1<<('a'-'a') | 1<<('c'-'a')
		assert.Equal(t, misa, conn.ReadCSR(csr.Misa))
		assert.Equal(t, 0b_100010001, conn.ReadCSR(csr.Mideleg))
		assert.Equal(t, 0, conn.ReadCSR(csr.Mie))
		assert.Equal(t, 1, conn.ReadCSR(csr.Mtvec))
		assert.Equal(t, 0, conn.ReadCSR(csr.Mscratch))
		assert.Equal(t, 0, conn.ReadCSR(csr.Mepc))
		assert.Equal(t, 0, conn.ReadCSR(csr.Mcause))
		assert.Equal(t, 0, conn.ReadCSR(csr.Mtval))
		assert.Equal(t, 0, conn.ReadCSR(csr.Mip))

		assert.Equal(t, 3<<csr.DcsrCause|3<<csr.DcsrPrv, conn.ReadCSR(csr.Dcsr))

		assert.Equal(t, 0x612, conn.ReadCSR(csr.Mvendorid))
		assert.Equal(t, 0x80000002, conn.ReadCSR(csr.Marchid))
		assert.Equal(t, 0x2, conn.ReadCSR(csr.Mimpid))
		assert.Equal(t, 0, conn.ReadCSR(csr.Mhartid))

		assert.Equal(t, 0x4000_0000, conn.ReadPC())

		for i := range 32 {
			assert.Equal(t, 0, conn.ReadGPR(i))
		}

		conn.WriteGPR(isa.A0, 0x12345678)
		assert.Equal(t, 0x12345678, conn.ReadGPR(isa.A0))
	})
}
