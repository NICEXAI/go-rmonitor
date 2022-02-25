//go:build windows
// +build windows

package cpu

import (
	"github.com/NICEXAI/go-rmonitor/util"
	"unsafe"
)

var (
	getActiveProcessorCount = util.Kernel32.NewProc("GetActiveProcessorCount")
	getNativeSystemInfo     = util.Kernel32.NewProc("GetNativeSystemInfo")
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
