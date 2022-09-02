package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
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

	df := dataframe.New(
		series.New([]string{"a", "b", "c", "d", "e"}, series.String, "alpha"),
		series.New([]int{5, 4, 3, 2, 1}, series.Int, "numbers"),
		series.New([]string{"a1", "b2", "c3", "d4", "e5"}, series.String, "alNums"),
		series.New([]bool{true, false, true, true, false}, series.Bool, "state"),
	)
	fmt.Println(df)

}
