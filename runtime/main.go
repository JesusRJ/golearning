package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func printStats(mem runtime.MemStats) {
	fmt.Println("---------------------------")
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("---------------------------")
}

// Print Go runtime version info
func printVersion() {
	myVersion := runtime.Version()
	major := strings.Split(myVersion, ".")[0][2]
	minor := strings.Split(myVersion, ".")[1]
	m1, _ := strconv.Atoi(string(major))
	m2, _ := strconv.Atoi(minor)
	if m1 == 1 && m2 < 8 {
		fmt.Println("Need Go version 1.8 or higher!")
		return
	}
	fmt.Printf("You are using Go version 1.8 or higher! (%s)\n", myVersion)
}

func main() {
	printVersion()

	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
	}

	printStats(mem)

	fmt.Println("NumCPU:", runtime.NumCPU())

	// for i := 0; i < 10; i++ {
	// 	s := make([]byte, 50000000)
	// 	if s == nil {
	// 		fmt.Println("Operation failed!")
	// 	}
	// 	time.Sleep(5 * time.Second)
	// }

	// printStats(mem)
}
