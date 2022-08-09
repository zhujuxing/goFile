package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	cpuNum := 8
	fmt.Println("Maximum CPu num:", runtime.NumCPU())
	n := runtime.GOMAXPROCS(cpuNum)
	fmt.Println("Pre-configure CPU num:", n)
	fmt.Println("After-configure CPU num:", cpuNum)

	last := time.Now()
	for i := 0; i < 100000; i++ {
		go func() {
			a := 999999 ^ 999999
			a = a + 1
		}()
	}
	now := time.Now()
	fmt.Println(now.Sub(last))
}
