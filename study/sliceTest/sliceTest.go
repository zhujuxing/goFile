package main

import "fmt"

func main() {
	//var numbers = make([]int, 3, 5)
	numbers := []int{0, 1, 2, 3}
	printSlice(numbers)

	numbers = append(numbers, 4)
	printSlice(numbers)

	fmt.Println(numbers[3])

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
