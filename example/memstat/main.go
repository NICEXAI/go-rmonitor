package main

import (
	"github.com/NICEXAI/go-rmonitor/mem"
	"log"
)

func main() {
	memStat, _ := mem.GetMemory()

	log.Printf("Totol memory is: %v", memStat.Total)
	log.Printf("Available memory is: %v", memStat.Available)
	log.Printf("Used memory is: %v", memStat.Used)
	log.Printf("Cached is: %v", memStat.Cached)
	log.Printf("UsedPercent is: %v", memStat.UsedPercent)
}
