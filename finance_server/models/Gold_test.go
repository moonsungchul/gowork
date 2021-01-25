package models_test

import (
	"testing"

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
}
