//       ___  _____  ____
//      / _ \/  _/ |/_/ /____ ______ _
//     / ___// /_>  </ __/ -_) __/  ' \
//    /_/  /___/_/|_|\__/\__/_/ /_/_/_/
//
//    Copyright 2019 Eliuk Blau
//
//    This Source Code Form is subject to the terms of the Mozilla Public
//    License, v. 2.0. If a copy of the MPL was not distributed with this
//    file, You can obtain one at https://mozilla.org/MPL/2.0/.

// build +windows

package main

import (
	"os"
	"syscall"
	"unsafe"
)

func init() {
	if isTerminal() {
		var consoleMode int32
		handle := int(os.Stdout.Fd())

		kernel32 := syscall.NewLazyDLL("kernel32.dll")
		procGetConsoleMode := kernel32.NewProc("GetConsoleMode")
		procSetConsoleMode := kernel32.NewProc("SetConsoleMode")

		// Enable VT100 escape sequences on Windows Console (disabled by default) - thanks @brutestack
		procGetConsoleMode.Call(uintptr(handle), uintptr(unsafe.Pointer(&consoleMode)))
		consoleMode |= 0x0004
		procSetConsoleMode.Call(uintptr(handle), uintptr(consoleMode))
	}
}
