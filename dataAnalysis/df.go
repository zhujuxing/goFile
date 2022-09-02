package main

import (
	"context"
	"fmt"
	"github.com/rocketlaunchr/dataframe-go"
	"github.com/rocketlaunchr/dataframe-go/utils/faker"
	"golang.org/x/exp/rand"
	"time"
)

func main() {
	s1 := dataframe.NewSeriesInt64("day", nil, 1, 2, 3, 4, 5, 6, 7, 8)
	s2 := dataframe.NewSeriesFloat64("sales", nil, 50.3, 23.4, 56.2, nil, nil, 84.2, 72, 89)
	df := dataframe.NewDataFrame(s1, s2)
	fmt.Println(df)

	df.Append(nil, 9, 123.6)

	df.Append(nil, map[string]interface{}{
		"day":   10,
		"sales": nil,
	})

	df.Remove(0)

	fmt.Println(df)

	sks := []dataframe.SortKey{
		{Key: "sales", Desc: true},
		{Key: "day", Desc: true},
	}

	ctx := context.Background()
	df.Sort(ctx, sks)

	fmt.Println(df)

	iter := df.ValuesIterator(dataframe.ValuesOptions{
		InitialRow:   0,
		Step:         1,
		DontReadLock: true,
	})

	df.Lock()
	for {
		row, vals, _ := iter()
		if row == nil {
			break
		}
		fmt.Println(*row, vals)
		//if vals["day"] == 9 {
		//	df.Update()
		//}
	}
	df.Unlock()

	fmt.Println(df)

	src := rand.NewSource(uint64(time.Now().UTC().UnixNano()))
	df1 := faker.NewDataFrame(8, src, faker.S("name", 0, "Name"), faker.S("title", 0.5, "JobTitle"), faker.S("base rate", 0, "Number", 15, 50))

	fmt.Println(df1)

	s := df1.Series[2]
	applyFn := dataframe.ApplySeriesFn(
		func(val interface{}, row int, nRows int) interface{} {
			return 2 * val.(int64)
		})
	_, err := dataframe.Apply(ctx, s, applyFn, dataframe.FilterOptions{InPlace: true})
	if err != nil {
		return
	}
	fmt.Println(df1)
}
