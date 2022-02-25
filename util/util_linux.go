//go:build linux
// +build linux

package util

import (
	"strconv"
	"strings"
)

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
