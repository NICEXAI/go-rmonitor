//go:build linux
// +build linux

package cpu

import (
	"github.com/NICEXAI/go-rmonitor/util"
	"strconv"
	"strings"
)

const (
	procCpuInfoHost = "/proc/cpuinfo"
	procStatHost    = "/proc/stat"
	sysCpuInfoHost  = "/sys/fs/cgroup/cpu/cpu.cfs_quota_us"
)

func GetCoreNum() float32 {
	ret := 0
	lines, err := util.ReadLines(procCpuInfoHost)
	if err == nil {
		for _, line := range lines {
			line = strings.ToLower(line)
			if strings.HasPrefix(line, "processor") {
				if _, err = strconv.Atoi(strings.TrimSpace(line[strings.IndexByte(line, ':')+1:])); err == nil {
					ret++
				}
			}
		}
	}
	if ret == 0 {
		lines, err = util.ReadLines(procStatHost)
		if err != nil {
			return -1
		}
		for _, line := range lines {
			if len(line) >= 4 && strings.HasPrefix(line, "cpu") && '0' <= line[3] && line[3] <= '9' { // `^cpu\d` regexp matching
				ret++
			}
		}
	}

	var (
		cpuInfo string
		cNum    int
	)
	cpuInfo, err = util.ReadContent(sysCpuInfoHost)
	if err != nil {
		return -1
	}

	cNum, err = strconv.Atoi(strings.ReplaceAll(cpuInfo, "\n", ""))
	if err != nil {
		return -1
	}
	if cNum == -1 {
		return float32(ret)
	}

	return float32(cNum / 100000)
}
