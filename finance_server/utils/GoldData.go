package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func main() {

	HighClose := func(s series.Series) series.Series {
		floats := s.Float()
		close := floats[4]
		high := floats[2]
		ch := close / high
		return series.Floats(ch)
	}

	OpenClose := func(s series.Series) series.Series {
		floats := s.Float()
		open := floats[1]
		close := floats[4]
		ch := open / close
		return series.Floats(ch)
	}

	LowClose := func(s series.Series) series.Series {
		floats := s.Float()
		low := floats[3]
		close := floats[4]
		ch := low / close
		return series.Floats(ch)
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
	//fmt.Println(df.Capply(mean))
	//fmt.Println(df.Rapply(close_high))
	closehigh := df.Rapply(HighClose)
	openclose := df.Rapply(OpenClose)
	lowclose := df.Rapply(LowClose)

	closehigh.SetNames("CloseHigh")
	openclose.SetNames("OpenClose")
	lowclose.SetNames("LowClose")
	df1 := df.CBind(closehigh)
	df2 := df1.CBind(openclose)
	df3 := df2.CBind(lowclose)
	fmt.Println(df3)
	fmt.Println(df3.Names())

	file, err := os.OpenFile("final.csv", os.O_CREATE, os.FileMode(0644))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	df3.WriteCSV(w)

}
