package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"log"
	"os"
)

func main() {
	f, e := os.Open("gotaTest/test.csv")
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(e)
		}
	}(f)
	if e != nil {
		log.Fatal(e)
	}

	df := dataframe.ReadCSV(f)
	fmt.Println(df)

	col1 := df.Subset([]int{0, 2})
	col2 := df.Select([]string{"Name", "Colour"})
	fmt.Println(col1)
	fmt.Println(col2)

	dataFrame2 := df.Set(
		[]int{0, 3},
		dataframe.LoadRecords(
			[][]string{
				{"Name", "Age", "Colour", "Height(ft)"},
				{"Jesse", "34", "indigo", "3.5"},
				{"Peter", "33", "violet", "3.3"},
			},
		),
	)

	fmt.Println(dataFrame2)

	// 更多gota的相关操作和例子，可以参考下面网址：
	//https://blog.csdn.net/weixin_32023091/article/details/112449908

	//df.Capply(
	//	func())

}
