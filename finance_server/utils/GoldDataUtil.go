package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

type GoldDataUtil struct {
	infile  string
	gold_df dataframe.DataFrame
}

func NewGoldDataUtil(file string) *GoldDataUtil {
	return &GoldDataUtil{infile: file}
}

func (r *GoldDataUtil) LoadCSV() {
	csvfile, err := os.Open(r.infile)
	if err != nil {
		log.Fatal(err)
	}
	r.gold_df = dataframe.ReadCSV(csvfile)
	fmt.Println(r.gold_df)
	r.CalData()

}

func (r *GoldDataUtil) CalData() {
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

	closehigh := r.gold_df.Rapply(HighClose)
	openclose := r.gold_df.Rapply(OpenClose)
	lowclose := r.gold_df.Rapply(LowClose)

	closehigh.SetNames("CloseHigh")
	openclose.SetNames("OpenClose")
	lowclose.SetNames("LowClose")

	r.gold_df = r.gold_df.CBind(closehigh)
	r.gold_df = r.gold_df.CBind(openclose)
	r.gold_df = r.gold_df.CBind(lowclose)
	fmt.Println(r.gold_df)

	r.gold_df = r.gold_df.CBind(r.CalCloseDivLastClose(5, "Close5LastClose"))
	r.gold_df = r.gold_df.CBind(r.CalCloseDivLastClose(10, "Close10LastClose"))
	r.gold_df = r.gold_df.CBind(r.CalCloseDivLastClose(20, "Close20LastClose"))
	r.gold_df = r.gold_df.CBind(r.CalCloseDivLastClose(60, "Close60LastClose"))
	r.gold_df = r.gold_df.CBind(r.CalCloseDivLastClose(120, "Close120LastClose"))
	r.gold_df = r.gold_df.CBind(r.CalVolumeDivLastVolume(5, "Volume5LastVolume"))
	r.gold_df = r.gold_df.CBind(r.CalVolumeDivLastVolume(10, "Volume10LastVolume"))
	r.gold_df = r.gold_df.CBind(r.CalVolumeDivLastVolume(20, "Volume20LastVolume"))
	r.gold_df = r.gold_df.CBind(r.CalVolumeDivLastVolume(60, "Volume60LastVolume"))
	r.gold_df = r.gold_df.CBind(r.CalVolumeDivLastVolume(120, "Volume120LastVolume"))
	fmt.Println("r.gold_df : ", r.gold_df)
}

func (r *GoldDataUtil) ToCSV(fname string) {
	f, err := os.Create(fname)
	if err != nil {
		fmt.Println(err)
		return
	}
	r.gold_df.WriteCSV(f)
}

/*
당일 종가 / 전일 종가
*/
func (r *GoldDataUtil) CalCloseLastClose() dataframe.DataFrame {
	ar := r.gold_df.Col("Close")
	items := ar.Float()
	var dt []float64 = make([]float64, len(items))

	before := items[0]
	for i := 1; i < len(items); i++ {
		fmt.Println(i)
		dt[i] = items[i] / before
		before = items[i]
	}
	df := dataframe.New(series.New(dt, series.Float, "CloseLastClose"))
	//se := series.New(dt, series.Float, "CloseLastClose")
	return df
}

/*
당일 거래량  /  전일 거래량
*/
func (r *GoldDataUtil) CalVolumeLastVolume() dataframe.DataFrame {
	ar := r.gold_df.Col("Volume")
	items := ar.Float()
	var dt []float64 = make([]float64, len(items))
	before := items[0]
	for i := 1; i < len(items); i++ {
		fmt.Println(i)
		dt[i] = items[i] / before
		before = items[i]
	}
	df := dataframe.New(series.New(dt, series.Float, "VolumeLastVolume"))
	fmt.Println("df : ", df)
	return df
}

/*
종가 / 5일 종가 평균
*/
func (r *GoldDataUtil) CalCloseDivLastClose(before_day int, title string) dataframe.DataFrame {
	ar := r.gold_df.Col("Close")
	items := ar.Float()
	var dt []float64 = make([]float64, len(items))
	for i := 0; i < before_day; i++ {
		dt[i] = 0
	}
	for i := before_day; i < len(items); i++ {
		avg := r.getBeforeAvg(items, i, before_day)
		dt[i] = items[i] / avg
	}
	df := dataframe.New(series.New(dt, series.Float, title))
	fmt.Println("df : ", df)
	return df
}

func (r *GoldDataUtil) CalVolumeDivLastVolume(before_day int, title string) dataframe.DataFrame {
	ar := r.gold_df.Col("Volume")
	items := ar.Float()
	var dt []float64 = make([]float64, len(items))
	for i := 0; i < before_day; i++ {
		dt[i] = 0
	}
	for i := before_day; i < len(items); i++ {
		avg := r.getBeforeAvg(items, i, before_day)
		dt[i] = items[i] / avg
	}
	df := dataframe.New(series.New(dt, series.Float, title))
	fmt.Println("df : ", df)
	return df
}

/*
지정된 날짜의 today_index 의 before_day의 평균을 구한다.
*/
func (r *GoldDataUtil) getBeforeAvg(items []float64, today_index int, before_day int) float64 {
	total := 0.0
	for i := today_index - before_day; i < today_index; i++ {
		total += items[i]
	}
	return total / float64(before_day)
}
