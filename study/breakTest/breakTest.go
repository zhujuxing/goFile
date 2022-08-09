package main

import "fmt"

func main() {
	i := 1
OuterLoop:
	for {
		for {
			if i > 5 {
				break OuterLoop //跳出OuterLoop标签对应的循环
			}
			fmt.Println(i)
			i++
		}
	}
}
