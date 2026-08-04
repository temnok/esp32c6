package tests

import (
	"runtime/debug"
	"testing"
)

func handlePanic(t *testing.T) {
	if err := recover(); err != nil {
		t.Fatalf("%v\n%s", err, debug.Stack())
	}
}
