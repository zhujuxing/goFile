package main

import (
	"fmt"
	"github.com/huandu/go-clone"
	"gonum.org/v1/gonum/graph/network"
	"log"
	"os"

	//"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/graphs/gen"
	"gonum.org/v1/gonum/graph/simple"
)

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
	//fmt.Printf("%s\n", b)

	file, err := os.Create("graph/gnp(10,0.6).dot")
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write(b)
	if err != nil {
		return nil
	}
	err = file.Close()
	if err != nil {
		return nil
	}
	return dst
}

func main() {
	//	gBuilder := graph.Builder
	//	gen.Gnp(gBuilder)
	//ExampleStar_undirectedRange()
	g := genGnpUndirected(10, 0.6)
	println(g.Nodes())
	//fmt.Println(g.From(99));
	gbc := network.Betweenness(g)
	//fmt.Println(gbc)
	traffic := clone.Clone(gbc).(map[int64]float64)
	//var capacity map[int64]floati64
	capacity := make(map[int64]float64)
	for k, v := range traffic {
		capacity[k] = (1 + 0.1) * v
	}
	//fmt.Println(gbc)
	//fmt.Println(traffic)
	fmt.Print(capacity) //打印介数字典

}
