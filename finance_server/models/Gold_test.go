package models_test

import (
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
	db.AutoMigrate(&models.UsbEur{})
	db.AutoMigrate(&models.UsbJpy{})
	db.AutoMigrate(&models.UsbKrw{})
	db.AutoMigrate(&models.UsbRub{})

	//infile := "../utils/gold_final.csv"
	store := models.MysqlStore{}
	db = store.Open("localhost", 3306, "fms_finance", "moonstar", "wooag01")
	//store.InsertGoldPrice(db, infile)
	store.InsertUSDaller(db, "C:/work/data_workspace/finance/data/usb_eur.csv", "USBEUR")
	store.InsertUSDaller(db, "C:/work/data_workspace/finance/data/usb_krw.csv", "USBKRW")
	store.InsertUSDaller(db, "C:/work/data_workspace/finance/data/usb_rub.csv", "USBRUB")
	store.InsertUSDaller(db, "C:/work/data_workspace/finance/data/usb_jpy.csv", "USBJPY")
	//store := models.MysqlStore{}
	/*
		all := store.GetPrices(db)

			for _, value := range allusers {
				fmt.Println(value)
			}

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
	*/

}

func Test2(t *testing.T) {
	store := models.MysqlStore{}
	db := store.Open("127.0.0.1", 3306, "fms_finance", "moonstar", "wooag01")

	all := store.GetPrices(db)
	count := store.GetPagesTotal(db)
	fmt.Println(count)
	fmt.Println(len(all))
	total := int64(len(all))
	assert.Equal(t, count, total)
}
