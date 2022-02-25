package main

import (
	"encoding/json"
	"fmt"
	"github.com/NICEXAI/go-rmonitor/cpu"
	"github.com/NICEXAI/go-rmonitor/mem"
	cpu2 "github.com/shirou/gopsutil/cpu"
	mem2 "github.com/shirou/gopsutil/mem"
)

func main() {
	count, _ := cpu2.Counts(true)
	fmt.Println(cpu.GetCoreNum(), count)

	memInfo, _ := mem.GetMemory()
	bMem, _ := json.Marshal(memInfo)

	memInfo2, _ := mem2.VirtualMemory()
	bMem2, _ := json.Marshal(memInfo2)
	fmt.Println(string(bMem))
	fmt.Println(string(bMem2))
}
