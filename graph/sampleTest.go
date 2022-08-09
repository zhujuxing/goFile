package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Sample(slice []int, n int) []interface{} {
	res := make([]interface{}, n)
	//order := make([]int,n)
	m := len(slice)
	order := rand.Perm(m)
	fmt.Println(order)
	for i := 0; i < n; i++ {
		res[i] = slice[order[i]]
		//println(i, " ", j)
	}
	return res
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//rand.Seed(time.Now().UnixNano())
	//for i := 0; i < 10; i++ {
	//	fmt.Println(rand.Perm(10))
	//}
	num := []int{1, 2, 3, 4, 5}
	fmt.Println(Sample(num, 2))
}
