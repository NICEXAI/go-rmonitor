//go:build windows
// +build windows

package util

import "golang.org/x/sys/windows"

var (
	Kernel32 = windows.NewLazySystemDLL("kernel32.dll")
)
