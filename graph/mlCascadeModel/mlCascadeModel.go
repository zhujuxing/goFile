package main

import (
	"fmt"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/graphs/gen"
	"gonum.org/v1/gonum/graph/network"
	"gonum.org/v1/gonum/graph/simple"
	"log"
	"math"
	"math/rand"
	"time"
)

func BetweennessNormal(g graph.Graph) map[int64]float64 {
	bc := network.Betweenness(g)
	n := len(bc)
	for i, j := range bc {
		bc[i] = j / float64((n-1)*(n-2))
	}
	return bc
}

func SampleIntSlice(slice []int, n int) []int {
	rand.Seed(time.Now().UnixNano())
	res := make([]int, n)
	m := len(slice)
	order := rand.Perm(m)
	for i := 0; i < n; i++ {
		res[i] = slice[order[i]]
	}
	return res
}

func main() {
	n := 100
	p := 0.3
	alpha := 0.1
	attackRate := 0.6
	monteCarloNum := 1000
	//monteCarloNum := 1000

	t1 := time.Now()

	g := simple.NewDirectedGraph()
	err := gen.Gnp(g, n, p, nil)
	if err != nil {
		log.Fatal(err)
	}

	attr := make(map[int64]map[string]float64)
	gbc := BetweennessNormal(g)

	for k := range gbc {
		var attrOfNode = map[string]float64{
			"load":     gbc[k],
			"capacity": (1 + alpha) * gbc[k],
		}
		attr[k] = attrOfNode
	}

	for i := 0; i < monteCarloNum; i++ {
		gIter := simple.NewDirectedGraph()
		//graph.Copy(g, gIter)
		graph.Copy(gIter, g)
		nSlice := make([]int, n)
		for j := 0; j < n; j++ {
			nSlice[j] = j
		}
		nodeFail := SampleIntSlice(nSlice, int(math.Ceil(float64(n)*attackRate)))
		for {
			if len(nodeFail) == 0 {
				break
			}
			for _, k := range nodeFail {
				gIter.RemoveNode(int64(k))
			}
			gbcIter := BetweennessNormal(gIter)
			nodeFail = make([]int, 0)
			for l := range gbcIter {
				attr[l]["load"] = gbcIter[l]

				if attr[l]["capacity"] < attr[l]["load"] {
					nodeFail = append(nodeFail, int(l))
				}
			}
		}
		fmt.Printf("第%d次仿真后，网络存活%d个节点\n", i, gIter.Nodes().Len())
	}
	t2 := time.Since(t1)
	fmt.Println("花费时间：", t2)
}
