//go:build linux
// +build linux

package util

import (
	"errors"
	"strconv"
	"strings"
)

const procCgroupHost = "/proc/1/cgroup"

var ErrInvalidFormat = errors.New("cgroups: parsing file with invalid format failed")

func GetContentFromCGroupFile(filename string) int {
	var (
		fileInfo string
		cNum     int
		err      error
	)
	fileInfo, err = ReadContent(filename)
	if err != nil {
		return -1
	}

	cNum, err = strconv.Atoi(strings.ReplaceAll(fileInfo, "\n", ""))
	if err != nil {
		return -1
	}
	return cNum
}

func ParseKV(raw string) (string, uint64, error) {
	parts := strings.Fields(raw)
	switch len(parts) {
	case 2:
		v, err := parseUint(parts[1], 10, 64)
		if err != nil {
			return "", 0, err
		}
		return parts[0], v, nil
	default:
		return "", 0, ErrInvalidFormat
	}
}

func parseUint(s string, base, bitSize int) (uint64, error) {
	v, err := strconv.ParseUint(s, base, bitSize)
	if err != nil {
		intValue, intErr := strconv.ParseInt(s, base, bitSize)
		// 1. Handle negative values greater than MinInt64 (and)
		// 2. Handle negative values lesser than MinInt64
		if intErr == nil && intValue < 0 {
			return 0, nil
		} else if intErr != nil &&
			intErr.(*strconv.NumError).Err == strconv.ErrRange &&
			intValue < 0 {
			return 0, nil
		}
		return 0, err
	}
	return v, nil
}

func IsContainer() bool {
	lines, err := ReadLines(procCgroupHost)
	if err != nil {
		return false
	}

	cpuStat := false
	memStat := false

	for _, line := range lines {
		lineArr := strings.Split(line, ":")
		if len(lineArr) != 3 {
			continue
		}
		fields := strings.Split(lineArr[1], ",")
		for _, field := range fields {
			if field == "cpu" && lineArr[2] != "/" {
				cpuStat = true
			}
			if field == "memory" && lineArr[2] != "/" {
				memStat = true
			}
		}
	}

	return cpuStat && memStat
}
