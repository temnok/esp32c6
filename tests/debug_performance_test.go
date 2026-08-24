package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/temnok/esp32c6/check"
	"github.com/temnok/esp32c6/debug"
	"github.com/temnok/esp32c6/isa/csr"
	"testing"
)

func TestPerformance(t *testing.T) {
	defer check.RecoverAndPrintStack(t.Fatal)

	debug.Session(func(conn *debug.Conn) {
		conn.HartResetAndHalt(0)

		assert.Equal(t, 0, conn.ReadCSR(csr.Mpcer))
		assert.Equal(t, 3, conn.ReadCSR(csr.Mpcmr))

		a := conn.ReadCSR(csr.Mpccr)
		b := conn.ReadCSR(csr.Mpccr)

		fmt.Println(b - a)
	})
}
