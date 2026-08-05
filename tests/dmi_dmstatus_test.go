package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/temnok/esp32c6/dmi"
	"testing"
)

func TestDmiDmstatus(t *testing.T) {
	defer handlePanic(t)

	dmi.Session(func(conn *dmi.Connection) {
		dmstatus := conn.Read(dmi.Dmstatus)
		assert.Equal(t, 2, dmstatus&0xF) // 2 (0.13): There is a Debug Module and it conforms to version 0.13 of this specification
	})
}
