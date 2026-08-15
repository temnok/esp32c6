package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/temnok/esp32c6/check"
	"github.com/temnok/esp32c6/debug"
	"testing"
)

func TestRegistersAfterReset(t *testing.T) {
	defer check.RecoverAndPrintStack(t.Fatal)

	debug.Session(func(conn *debug.Conn) {
		conn.HartResetAndHalt(0)

		assert.Equal(t, 0, conn.ReadCSR(0x000)) // ustatus
		assert.Equal(t, 0, conn.ReadCSR(0x004)) // uie
		assert.Equal(t, 1, conn.ReadCSR(0x005)) // utvec
		//		assert.Equal(t, 0, conn.ReadCSR(0x040)) // uscratch ???
		assert.Equal(t, 0, conn.ReadCSR(0x041)) // uepc
		assert.Equal(t, 0, conn.ReadCSR(0x042)) // ucause
		assert.Equal(t, 0, conn.ReadCSR(0x043)) // utval
		assert.Equal(t, 0, conn.ReadCSR(0x044)) // uip

		assert.Equal(t, 1<<21|3<<11, conn.ReadCSR(0x300)) // mstatus: TW | MPP
		misa := 1<<30 | 1<<('n'-'a') | 1<<('u'-'a') | 1<<('x'-'a') |
			1<<('i'-'a') | 1<<('m'-'a') | 1<<('a'-'a') | 1<<('c'-'a')
		assert.Equal(t, misa, conn.ReadCSR(0x301))
		assert.Equal(t, 0b_100010001, conn.ReadCSR(0x303)) // mideleg
		assert.Equal(t, 0, conn.ReadCSR(0x304))            // mie
		assert.Equal(t, 1, conn.ReadCSR(0x305))            // mtvec
		assert.Equal(t, 0, conn.ReadCSR(0x340))            // mscratch
		assert.Equal(t, 0, conn.ReadCSR(0x341))            // mepc
		assert.Equal(t, 0, conn.ReadCSR(0x342))            // mcause
		assert.Equal(t, 0, conn.ReadCSR(0x343))            // mtval
		assert.Equal(t, 0, conn.ReadCSR(0x344))            // mip

		assert.Equal(t, 0x612, conn.ReadCSR(0xF11))      // vendorid
		assert.Equal(t, 0x80000002, conn.ReadCSR(0xF12)) // marchid
		assert.Equal(t, 0x2, conn.ReadCSR(0xF13))        // mimpid
		assert.Equal(t, 0, conn.ReadCSR(0xF14))          // mhartid

		assert.Equal(t, 0x4000_0000, conn.ReadPC())

		for i := range 32 {
			assert.Equal(t, 0, conn.ReadGPR(i))
		}
	})
}
