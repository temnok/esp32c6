package tests

import (
	"github.com/google/gousb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJtagEndpoint(t *testing.T) {
	defer handlePanic(t)

	withJtagUsbEndpoints(func(in *gousb.InEndpoint, out *gousb.OutEndpoint) {
		assert.Equal(t, "ep #3 IN (address 0x83) bulk [64 bytes]", in.String())
		assert.Equal(t, "ep #2 OUT (address 0x02) bulk [64 bytes]", out.String())
	})
}
