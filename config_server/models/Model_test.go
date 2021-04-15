package models_test

import (
	"testing"

	"github.com/moonsungchul/configserver/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test1(t *testing.T) {
	dsn := "moonstar:wooag01@tcp(172.17.0.3:3306)/fms_config?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Config{})
	db.AutoMigrate(&models.ConfigSection{})
	db.AutoMigrate(&models.ConfigProperty{})

}
