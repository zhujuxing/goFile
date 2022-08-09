package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"strings"
)

func loadFromCSV(csv string) dataframe.DataFrame {
	df := dataframe.ReadCSV(strings.NewReader(csv), dataframe.WithDelimiter(rune(',')))
	//println(df)
	return df
}

func main() {
	csvStr := `
Country,Date,Age,Amount,Id
"United States",2012-02-01,50,112.1,01234
"United States",2012-02-01,32,321.31,54320
"United Kingdom",2012-02-01,17,18.2,12345
"United States",2012-02-01,32,321.31,54320
"United Kingdom",2012-05-07,NA,18.2,12345
"United States",2012-02-01,32,321.31,54320
"United States",2012-02-01,32,321.31,54320
Spain,2012-02-01,66,555.42,00241
`
	//rd := io.Reader("128server业务可用度计算结果T1000.csv")
	//csv.NewReader()
	df := loadFromCSV(csvStr)
	fmt.Println(df)
}
