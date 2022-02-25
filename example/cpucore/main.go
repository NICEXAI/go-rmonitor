package main

import (
	"github.com/NICEXAI/go-rmonitor/cpu"
	"log"
)

func main() {
	coreCount := cpu.GetCoreNum()
	log.Printf("CPU core count is: %v", coreCount)
}
