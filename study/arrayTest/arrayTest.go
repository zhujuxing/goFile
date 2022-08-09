package main

import "fmt"

func main() {
	var student = [...]string{"Tom", "Ben", "Peter"}
	fmt.Println(student)
	for k, v := range student {
		fmt.Println("Key is :", k, " Value is :", v)
	}
}
