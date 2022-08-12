package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func main() {
	df := dataframe.New(
		series.New([]string{"a", "b", "c", "d", "e"}, series.String, "alpha"),
		series.New([]int{5, 4, 3, 2, 1}, series.Int, "numbers"),
		series.New([]string{"a1", "b2", "c3", "d4", "e5"}, series.String, "alnums"),
		series.New([]bool{true, false, true, true, false}, series.Bool, "state"),
	)
	fmt.Println(df)

	//g := simple.NewDirectedGraph()
	//err := gen.Gnp(g, 10, 0.4, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//gdf := dataframe.LoadStructs(g)
	//fmt.Println(gdf)

}
