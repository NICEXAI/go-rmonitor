# go-rmonitor
An easy-to-use go service resource monitor, support containers.

### Installation

Run the following command under your project:

> go get -u github.com/NICEXAI/go-rmonitor

### Available Architectures

* Linux i386/amd64/arm(raspberry pi)
* Windows i386/amd64/arm/arm64

### Usage
##### Get CPU cores
```go
package main

import (
	"github.com/NICEXAI/go-rmonitor/cpu"
	"log"
)

func main() {
	coreCount := cpu.GetCoreNum()
	log.Printf("CPU core count is: %v", coreCount)
}
```
See [details](./example/cpucore/main.go).

##### Get memory stat
```go
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

```
See [details](./example/memstat/main.go).