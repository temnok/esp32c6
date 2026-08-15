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

		assert.Equal(t, 0x612, conn.ReadCSR(0xF11))      // vendorid
		assert.Equal(t, 0x80000002, conn.ReadCSR(0xF12)) // marchid
		assert.Equal(t, 0x2, conn.ReadCSR(0xF13))        // mimpid
		assert.Equal(t, 0, conn.ReadCSR(0xF14))          // mhartid

		expectedMisa := 1<<30 | 1<<('n'-'a') | 1<<('u'-'a') | 1<<('x'-'a') |
			1<<('i'-'a') | 1<<('m'-'a') | 1<<('a'-'a') | 1<<('c'-'a')
		assert.Equal(t, expectedMisa, conn.ReadCSR(0x301)) // misa

		assert.Equal(t, 0x40000000, conn.ReadPC())

		for i := range 32 {
			assert.Equal(t, 0, conn.ReadGPR(i))
		}
	})
}
