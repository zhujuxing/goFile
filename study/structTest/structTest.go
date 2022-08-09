package main

import "fmt"

type BookBorrow struct {
	Book struct {
		title  string
		author string
		num    int
		id     int
	}
	borrowTime string
}

func main() {
	bookBorrow := &BookBorrow{
		Book: struct {
			title  string
			author string
			num    int
			id     int
		}{
			"GO语言",
			"Tom",
			20,
			122368,
		},
		borrowTime: "30",
	}
	fmt.Println(bookBorrow)
}
