package main

import (
	"fmt"
	"github.com/huandu/go-clone"
	"gonum.org/v1/gonum/graph/network"
	"log"

	//"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/graphs/gen"
	"gonum.org/v1/gonum/graph/simple"
)

//func ExampleStar_undirectedRange() {
//	dst := simple.NewUndirectedGraph()
//	err := gen.Star(dst, 0, gen.IDRange{First: 1, Last: 6})
//	if err != nil {
//		log.Fatal(err)
//	}
//	b, err := dot.Marshal(dst, "star", "", "\t")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("%s\n", b)
//}

func genGnpUndirected(n int, p float64) *simple.DirectedGraph {
	dst := simple.NewDirectedGraph()
	err := gen.Gnp(dst, n, p, nil)
	if err != nil {
		log.Fatal(err)
	}
	b, err := dot.Marshal(dst, "star", "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", b)
	return dst
}
func main() {

	//	gBuilder := graph.Builder
	//	gen.Gnp(gBuilder)
	//ExampleStar_undirectedRange()
	g := genGnpUndirected(100, 0.6)
	//fmt.Println(g.From(99));
	gbc := network.Betweenness(g)
	//fmt.Println(gbc)
	traffic := clone.Clone(gbc).(map[int64]float64)
	var capacity map[int64]float64
	for k, v := range traffic {
		capacity[k] = (1 + 0.1) * v
	}
	fmt.Print(capacity)

}
