package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func main() {
	for i, v := range make([]string, 10) {
		once.Do(f)
		fmt.Println("v:", v, "--i:", i)
	}

	for i := 0; i < 10; i++ {
		go func(i int) {
			once.Do(fDone)
			fmt.Println(i)
		}(i)
	}
	time.Sleep(10)
}

func f() {
	fmt.Println("once")
}

func fDone() {
	fmt.Println("onced")
}
