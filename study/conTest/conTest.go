package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c = "what"
		d
		e = iota
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
