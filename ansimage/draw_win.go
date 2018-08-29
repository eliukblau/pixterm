// +build windows

package ansimage

import (
	"syscall"
	"unsafe"

	"golang.org/x/crypto/ssh/terminal"
)

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	procGetConsoleMode = kernel32.NewProc("GetConsoleMode")
	procSetConsoleMode = kernel32.NewProc("SetConsoleMode")
)

func EnableColorsInTerminal() {
	handle := syscall.Stdout
	if terminal.IsTerminal(int(handle)) {

		var consoleMode int32

		procGetConsoleMode.Call(uintptr(handle), uintptr(unsafe.Pointer(&consoleMode)))
		consoleMode |= 0x0004
		procSetConsoleMode.Call(uintptr(handle), uintptr(consoleMode))
	}
}
