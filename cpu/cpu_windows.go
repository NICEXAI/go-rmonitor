//go:build windows
// +build windows

package cpu

import (
	"golang.org/x/sys/windows"
	"unsafe"
)

var (
	kernel32                = windows.NewLazySystemDLL("kernel32.dll")
	getActiveProcessorCount = kernel32.NewProc("GetActiveProcessorCount")
	getNativeSystemInfo     = kernel32.NewProc("GetNativeSystemInfo")
)

type systemInfo struct {
	wProcessorArchitecture      uint16
	wReserved                   uint16
	dwPageSize                  uint32
	lpMinimumApplicationAddress uintptr
	lpMaximumApplicationAddress uintptr
	dwActiveProcessorMask       uintptr
	dwNumberOfProcessors        uint32
	dwProcessorType             uint32
	dwAllocationGranularity     uint32
	wProcessorLevel             uint16
	wProcessorRevision          uint16
}

// GetCoreNum returns the number of cpu cores in the system
func GetCoreNum() float32 {
	if err := getActiveProcessorCount.Find(); err == nil {
		if ret, _, _ := getActiveProcessorCount.Call(uintptr(0xffff)); ret != 0 {
			return float32(ret)
		}
	}

	var systemInfo systemInfo
	if _, _, err := getNativeSystemInfo.Call(uintptr(unsafe.Pointer(&systemInfo))); err != nil || systemInfo.dwNumberOfProcessors == 0 {
		return -1
	}
	return float32(systemInfo.dwNumberOfProcessors)
}
