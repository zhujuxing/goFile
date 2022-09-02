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
	"runtime"
	"sync"
	"time"
)

var lock sync.RWMutex

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

	cpuNum := 8
	fmt.Println("Maximum CPU num:", runtime.NumCPU())
	cpuNumNew := runtime.GOMAXPROCS(cpuNum)
	fmt.Println("Pre-configure CPU num:", cpuNumNew)
	fmt.Println("After-configure CPU num:", cpuNum)

	//t1 := time.Now()

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

	printAfterCalAll(monteCarloNum, g, n, attackRate, attr)
	printWhileCal(monteCarloNum, g, n, attackRate, attr)
}

func printWhileCal(monteCarloNum int, g *simple.DirectedGraph, n int, attackRate float64, attr map[int64]map[string]float64) {
	t1 := time.Now()
	res := make(chan int, monteCarloNum)
	for i := 0; i < monteCarloNum; i++ {
		go CascadeFailure(g, n, attackRate, attr, res)
		go func(c <-chan int) {
			fmt.Printf("第次仿真后，网络存在%d个节点\n", <-res)
		}(res)
	}

	t2 := time.Since(t1)
	fmt.Println("花费时间：", t2)
}

func printAfterCalAll(monteCarloNum int, g *simple.DirectedGraph, n int, attackRate float64, attr map[int64]map[string]float64) {
	t1 := time.Now()
	//res := make(map[int]int)
	res := make([]chan int, monteCarloNum)
	for i := 0; i < monteCarloNum; i++ {
		res[i] = make(chan int)
		go CascadeFailure(g, n, attackRate, attr, res[i])
	}

	//defer func(r map[int]int) {
	//	for i, j := range r {
	//		fmt.Printf("第%d次仿真后，网络存在%d个节点\n", i, j)
	//	}
	//}(res)
	for i, r := range res {
		fmt.Printf("第%d次仿真后，网络存在%d个节点\n", i, <-r)
	}
	t2 := time.Since(t1)
	fmt.Println("花费时间：", t2)
}

func CascadeFailure(g *simple.DirectedGraph, n int, attackRate float64, attr map[int64]map[string]float64, res chan int) {
	gIter := simple.NewDirectedGraph()
	lock.Lock()
	graph.Copy(gIter, g)
	lock.Unlock()
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
	//fmt.Printf("第%d次仿真后，网络存活%d个节点\n", i, gIter.Nodes().Len())
	res <- gIter.Nodes().Len()
}
