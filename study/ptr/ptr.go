package main

import "golang.org/x/exp/errors/fmt"

func main() {
	n := 1
	p := &n
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(&p)
}
