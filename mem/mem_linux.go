//go:build linux
// +build linux

package mem

import (
	"github.com/NICEXAI/go-rmonitor/util"
	"strconv"
	"strings"
)

const (
	procMemInfoHost      = "/proc/meminfo"
	sysMemLimitHost      = "/sys/fs/cgroup/memory/memory.limit_in_bytes"
	sysMemUsageLimitHost = "/sys/fs/cgroup/memory/memory.usage_in_bytes"
)

func GetMemory() (*MemoryStat, error) {
	lines, _ := util.ReadLines(procMemInfoHost)

	// flag if MemAvailable is in /proc/meminfo (kernel 3.14+)
	memAvail := false
	ret := &MemoryStat{}

	for _, line := range lines {
		fields := strings.Split(line, ":")
		if len(fields) != 2 {
			continue
		}
		key := strings.TrimSpace(fields[0])
		value := strings.TrimSpace(fields[1])
		value = strings.Replace(value, " kB", "", -1)

		switch key {
		case "MemTotal":
			t, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return ret, err
			}
			ret.Total = t * 1024
		case "MemFree":
			t, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return ret, err
			}
			ret.free = t * 1024
		case "MemAvailable":
			t, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return ret, err
			}
			memAvail = true
			ret.Available = t * 1024
		case "Buffers":
			t, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return ret, err
			}
			ret.Buffers = t * 1024
		case "Cached":
			t, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return ret, err
			}
			ret.Cached = t * 1024
		case "SReclaimable":
			t, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return ret, err
			}
			ret.sreClaimAble = t * 1024
		}
	}

	ret.Cached += ret.sreClaimAble

	if !memAvail {
		ret.Available = ret.Cached + ret.free
	}

	ret.Used = ret.Total - ret.free - ret.Buffers - ret.Cached

	//detect cgroup
	memLimit := util.GetContentFromCGroupFile(sysMemLimitHost)
	if memLimit != -1 && memLimit != 9223372036854771712 {
		ret.Total = uint64(memLimit)
	}

	memUsageLimit := util.GetContentFromCGroupFile(sysMemUsageLimitHost)
	if memUsageLimit != -1 {
		ret.Used = uint64(memUsageLimit)
	}

	ret.Available = ret.Total - ret.Used
	ret.UsedPercent = float64(ret.Used) / float64(ret.Total) * 100.0

	return ret, nil
}
