package models

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlStore struct {
}

func (r MysqlStore) Open(host string, port int16, dbname string,
	user string, passwd string) *gorm.DB {
	ss := fmt.Sprintf("%s:%d", host, port)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, passwd, ss, dbname)
	//dsn := "moonstar:wooag01@tcp(127.0.0.1:3306)/fms_finance?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(GoldPrice{})
	return db
}

func (r MysqlStore) GetPrices(db *gorm.DB) []GoldPrice {
	all := []GoldPrice{}
	db.Find(&all)
	return all
}

func (r MysqlStore) GetPricesPages(db *gorm.DB, start int, rows int) []GoldPrice {
	all := []GoldPrice{}
	db.Limit(rows).Offset(start).Find(&all)
	return all
}

func (r MysqlStore) GetPagesTotal(db *gorm.DB) int64 {
	var count int64
	db.Model(&GoldPrice{}).Count(&count)
	return count
}

func (r MysqlStore) InsertGoldPrice(db *gorm.DB, infile string) {
	content, err := ioutil.ReadFile(infile)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	for i, value := range lines[1:] {
		ar := strings.Split(strings.TrimSpace(value), ",")

		gopen, _ := strconv.ParseFloat(ar[1], 32)
		ghigh, _ := strconv.ParseFloat(ar[1], 32)
		glow, _ := strconv.ParseFloat(ar[1], 32)
		gclose, _ := strconv.ParseFloat(ar[1], 32)
		gvolume, _ := strconv.ParseFloat(ar[1], 32)

		md := GoldPrice{
			GDate:   ar[0],
			GOpen:   float32(gopen),
			GHigh:   float32(ghigh),
			GLow:    float32(glow),
			GClose:  float32(gclose),
			GVolume: float32(gvolume),
		}
		fmt.Println(i, md)
		db.Create(&md)
	}

}
