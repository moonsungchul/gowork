package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func main() {

	mean := func(s series.Series) series.Series {
		floats := s.Float()
		fmt.Println(floats)
		//for _, f := range floats {
		//	fmt.Println(f)
		//}
		return series.Floats(floats)
	}

	csvfile, err := os.Open("C:/work/data_workspace/finance/src/gold.csv")
	if err != nil {
		log.Fatal(err)
	}

	df := dataframe.ReadCSV(csvfile)
	fmt.Println(df)
	sub := df.Subset([]int{0, 1})
	fmt.Println(sub)

	//col := df.Col("Date")
	//fmt.Println(col)
	//close := df.Col("Close")
	//fmt.Println(close)
	//close := df.Col("Close")
	//fmt.Println(close)
	//df.Capply(mean)
	df.Rapply(mean)

}
