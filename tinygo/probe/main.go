package main

import (
	"unsafe"
)

//go:noinline
func main() {
	addr := uintptr(0x4087_0000)
	ptr := (*uint32)(unsafe.Pointer(addr))

	*ptr = 0
}
