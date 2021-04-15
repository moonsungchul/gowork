package models

import (
	"fmt"

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

	return db
}

func (r MysqlStore) GetConfigCount(db *gorm.DB, cname string) int64 {
	var count int64
	db.Model(&Config{CName: cname}).Count(&count)
	return count
}

func (r MysqlStore) InsertConfig(db *gorm.DB, cname string) {
	cc := Config{CName: cname}
	co := r.GetConfigCount(db, cname)
	if co > 0 {
		db.Create(&cc)
	}
}

func (r MysqlStore) UpdateConfig(db *gorm.DB, config Config) {
	db.Save(&config)
}

func (r MysqlStore) DeleteConfig(db *gorm.DB, config Config) {
	db.Delete(&config)
}

func (r MysqlStore) GetSectionCount(db *gorm.DB, config_id uint, section string) int64 {
	var count int64
	db.Model(&ConfigSection{ConfigID: config_id, Section: section}).Count(&count)
	return count
}

func (r MysqlStore) InsertSection(db *gorm.DB, config_id uint, section string) {
	cc := ConfigSection{ConfigID: config_id, Section: section}
	co := r.GetSectionCount(db, config_id, section)
	if co > 0 {
		db.Create(&cc)
	}
}

func (r MysqlStore) UpdateSection(db *gorm.DB, sec ConfigSection) {
	db.Save(&sec)
}

func (r MysqlStore) DeleteSection(db *gorm.DB, sec ConfigSection) {
	db.Delete(&sec)
}

func (r MysqlStore) GetPropertyCount(db *gorm.DB, config_id uint,
	sec_id uint, key string) int64 {
	var count int64
	db.Model(&ConfigProperty{ConfigID: config_id, SectionID: sec_id, Key: key}).Count(&count)
	return count
}

func (r MysqlStore) InsertPropery(db *gorm.DB, pro ConfigProperty) {
	co := r.GetPropertyCount(db, pro.ConfigID, pro.SectionID, pro.Key)
	if co > 0 {
		db.Create(&pro)
	}
}

func (r MysqlStore) UpdateProperty(db *gorm.DB, pro ConfigProperty) {
	db.Save(&pro)
}

func (r MysqlStore) DeleteProperty(db *gorm.DB, pro ConfigProperty) {
	db.Delete(&pro)
}
