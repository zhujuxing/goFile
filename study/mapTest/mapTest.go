package main

import (
	"fmt"
	"sync"
)

var lock sync.RWMutex

func main() {
	syncMap()
	mapInit()
}

func mapInit() {
	m := make(map[int]int)
	m[0] = 0
	m[1] = 1
	fmt.Println(m)
}

func syncMap() {
	GoMap := make(map[int]int)
	for i := 0; i < 100000; i++ {
		go writeMap(GoMap, i, i)
		go readMap(GoMap, i)
	}
	fmt.Println("Done")
}

func readMap(GoMap map[int]int, key int) int {
	lock.Lock()
	m := GoMap[key]
	lock.Unlock()
	return m
}

func writeMap(GoMap map[int]int, key int, value int) {
	lock.Lock()
	GoMap[key] = value
	lock.Unlock()
}
