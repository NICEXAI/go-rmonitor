package main

import (
	"fmt"
	"github.com/NICEXAI/go-rmonitor/cpu"
	cpu2 "github.com/shirou/gopsutil/cpu"
)

func main() {
	count, _ := cpu2.Counts(true)
	fmt.Println(cpu.GetCoreNum(), count)
	fmt.Println(float32(50000) / 100000)
}
