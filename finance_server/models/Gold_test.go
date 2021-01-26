package models_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/moonsungchul/finance/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test1(t *testing.T) {
	dsn := "moonstar:wooag01@tcp(127.0.0.1:3306)/fms_finance?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.GoldPrice{})

	//infile := "C:/work/data_workspace/finance/src/gold.csv"
	//kstore := models.MysqlStore{}
	//store.InsertGoldPrice(infile, db)
	store := models.MysqlStore{}
	all := store.GetPrices(db)

	/*
		for _, value := range allusers {
			fmt.Println(value)
		}
	*/

	count := store.GetPagesTotal(db)
	fmt.Println(count)
	fmt.Println(len(all))
	total := int64(len(all))
	assert.Equal(t, count, total)

	res := store.GetPricesPages(db, 10, 100)
	for _, value := range res {
		jj, err := json.Marshal(&value)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(jj))
	}
	assert.Equal(t, len(res), 100)

}
