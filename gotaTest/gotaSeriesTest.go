package main

import (
	"fmt"
	"github.com/go-gota/gota/series"
)

func main() {
	s1 := series.New([]string{"z", "y", "d", "e"}, series.String, "col")
	s2 := map[string]series.Type{
		"A": series.String,
		"D": series.Bool,
	}
	fmt.Println(s1)
	fmt.Println(s2)
}
