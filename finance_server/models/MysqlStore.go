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

func (r MysqlStore) InsertUSDaller(db *gorm.DB, infile string, stype string) {
	content, err := ioutil.ReadFile(infile)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	for _, value := range lines[1:] {
		ar := strings.Split(strings.TrimSpace(value), ",")

		gclose, _ := strconv.ParseFloat(ar[1], 32)
		gopen, _ := strconv.ParseFloat(ar[2], 32)
		ghigh, _ := strconv.ParseFloat(ar[3], 32)
		glow, _ := strconv.ParseFloat(ar[4], 32)
		gchange, _ := strconv.ParseFloat(ar[5], 32)

		if stype == "USKRW" {
			md := UsbKrw{
				GDate:   ar[0],
				GClose:  float32(gclose),
				GOpen:   float32(gopen),
				GHigh:   float32(ghigh),
				GLow:    float32(glow),
				GChange: float32(gchange),
			}
			db.Create(&md)
		} else if stype == "USBEUR" {
			md := UsbEur{
				GDate:   ar[0],
				GClose:  float32(gclose),
				GOpen:   float32(gopen),
				GHigh:   float32(ghigh),
				GLow:    float32(glow),
				GChange: float32(gchange),
			}
			db.Create(&md)
		} else if stype == "USBJPY" {
			md := UsbJpy{
				GDate:   ar[0],
				GClose:  float32(gclose),
				GOpen:   float32(gopen),
				GHigh:   float32(ghigh),
				GLow:    float32(glow),
				GChange: float32(gchange),
			}
			db.Create(&md)
		} else if stype == "USBRUB" {
			md := UsbRub{
				GDate:   ar[0],
				GClose:  float32(gclose),
				GOpen:   float32(gopen),
				GHigh:   float32(ghigh),
				GLow:    float32(glow),
				GChange: float32(gchange),
			}
			db.Create(&md)
		}
		//fmt.Println(i, md)
	}

}

func (r MysqlStore) InsertGoldPrice(db *gorm.DB, infile string) {
	content, err := ioutil.ReadFile(infile)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	for _, value := range lines[1:] {
		ar := strings.Split(strings.TrimSpace(value), ",")
		fmt.Println("ar len : ", len(ar))

		gopen, _ := strconv.ParseFloat(ar[1], 32)
		ghigh, _ := strconv.ParseFloat(ar[2], 32)
		glow, _ := strconv.ParseFloat(ar[3], 32)
		gclose, _ := strconv.ParseFloat(ar[4], 32)
		gvolume, _ := strconv.ParseFloat(ar[5], 32)

		closeHigh, _ := strconv.ParseFloat(ar[6], 32)
		openClose, _ := strconv.ParseFloat(ar[7], 32)
		lowClose, _ := strconv.ParseFloat(ar[8], 32)
		close5LastClose, _ := strconv.ParseFloat(ar[9], 32)
		close10LastClose, _ := strconv.ParseFloat(ar[10], 32)
		close20LastClose, _ := strconv.ParseFloat(ar[11], 32)
		close60LastClose, _ := strconv.ParseFloat(ar[12], 32)
		close120LastClose, _ := strconv.ParseFloat(ar[13], 32)

		volume5LastVolume, _ := strconv.ParseFloat(ar[14], 32)
		volume10LastVolume, _ := strconv.ParseFloat(ar[15], 32)
		volume20LastVolume, _ := strconv.ParseFloat(ar[16], 32)
		volume60LastVolume, _ := strconv.ParseFloat(ar[17], 32)
		volume120LastVolume, _ := strconv.ParseFloat(ar[18], 32)

		md := GoldPrice{
			GDate:               ar[0],
			GOpen:               float32(gopen),
			GHigh:               float32(ghigh),
			GLow:                float32(glow),
			GClose:              float32(gclose),
			GVolume:             float32(gvolume),
			CloseHigh:           float32(closeHigh),
			OpenClose:           float32(openClose),
			LowClose:            float32(lowClose),
			Close5LastClose:     float32(close5LastClose),
			Close10LastClose:    float32(close10LastClose),
			Close20LastClose:    float32(close20LastClose),
			Close60LastClose:    float32(close60LastClose),
			Close120LastClose:   float32(close120LastClose),
			Volume5LastVolume:   float32(volume5LastVolume),
			Volume10LastVolume:  float32(volume10LastVolume),
			Volume20LastVolume:  float32(volume20LastVolume),
			Volume60LastVolume:  float32(volume60LastVolume),
			Volume120LastVolume: float32(volume120LastVolume),
		}
		//fmt.Println(i, md)
		db.Create(&md)
	}

}
