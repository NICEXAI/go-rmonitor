//go:build windows
// +build windows

package mem

import (
	"github.com/NICEXAI/go-rmonitor/util"
	"golang.org/x/sys/windows"
	"unsafe"
)

var (
	globalMemoryStatusEx = util.Kernel32.NewProc("GlobalMemoryStatusEx")
)

type memoryStatusEx struct {
	cbSize                  uint32
	dwMemoryLoad            uint32
	ullTotalPhys            uint64 // in bytes
	ullAvailPhys            uint64
	ullTotalPageFile        uint64
	ullAvailPageFile        uint64
	ullTotalVirtual         uint64
	ullAvailVirtual         uint64
	ullAvailExtendedVirtual uint64
}

func GetMemory() (*MemoryStat, error) {
	var memInfo memoryStatusEx

	memInfo.cbSize = uint32(unsafe.Sizeof(memInfo))
	mem, _, _ := globalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&memInfo)))
	if mem == 0 {
		return nil, windows.GetLastError()
	}

	ret := &MemoryStat{
		Total:       memInfo.ullTotalPhys,
		Available:   memInfo.ullAvailPhys,
		UsedPercent: float64(memInfo.dwMemoryLoad),
	}

	ret.Used = ret.Total - ret.Available
	return ret, nil
}
