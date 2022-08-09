package main

import "fmt"

func main() {
	//var myInt int
	//var myIntPointer *int
	//myIntPointer = & myInt
	//fmt.Println(myIntPointer)

	myInt := 4.5
	myIntPointer := &myInt
	fmt.Println(myInt)
	*myIntPointer = 3
	fmt.Println(myInt)
}
