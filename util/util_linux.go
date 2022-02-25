//go:build linux
// +build linux

package util

import (
	"errors"
	"strconv"
	"strings"
)

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
