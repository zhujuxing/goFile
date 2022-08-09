package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("time包中的常量Second: %v \n", time.Second)
	fmt.Printf("time包中的UTC变量: %v \n", time.UTC)
	time.UTC, _ = time.LoadLocation("Local")
	fmt.Printf("time包中的UTC变量更改后的值： %v \n", time.UTC)
}
