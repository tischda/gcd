package main

import (
	"syscall"
)

var (
	modkernel32      = syscall.NewLazyDLL("kernel32.dll")
	procGetConsoleCP = modkernel32.NewProc("GetConsoleCP")
)

// https://golang.org/src/internal/syscall/windows/zsyscall_windows.go
func GetConsoleCP() (ccp uint32) {
	r0, _, _ := syscall.Syscall(procGetConsoleCP.Addr(), 0, 0, 0, 0)
	ccp = uint32(r0)
	return
}
