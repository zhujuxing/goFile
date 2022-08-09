package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"log"
	"os"
)

func main() {
	file, err := os.Open("gotaTest/128server业务可用度计算结果T1000.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	df := dataframe.ReadCSV(file)

	fmt.Println(df)
}
