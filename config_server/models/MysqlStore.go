package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlStore struct {
	Host   string
	Port   int16
	Dbname string
	User   string
	Passwd string
}

func (r MysqlStore) OpenDB() *gorm.DB {
	return r.Open(r.Host, r.Port, r.Dbname, r.User, r.Passwd)
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

func (r MysqlStore) CreateTable(db *gorm.DB) {
	db.AutoMigrate(&Config{})
	db.AutoMigrate(&ConfigSection{})
	db.AutoMigrate(&ConfigProperty{})
}

func (r MysqlStore) GetConfig(db *gorm.DB, cname string) Config {
	var cc Config
	db.Where("c_name = ?", cname).First(&cc)
	return cc
}

func (r MysqlStore) InsertConfig(db *gorm.DB, cname string) {
	cc := r.GetConfig(db, cname)
	if cc.CName == "" {
		conf := Config{CName: cname}
		db.Create(&conf)
	}
}

func (r MysqlStore) UpdateConfig(db *gorm.DB, config Config) {
	db.Save(&config)
}

func (r MysqlStore) DeleteConfig(db *gorm.DB, config Config) {
	db.Delete(&Config{}, config.ID)
}

func (r MysqlStore) GetSection(db *gorm.DB, config_id uint, section string) ConfigSection {
	var sec ConfigSection
	db.Where("section = ? and config_id = ?", section, config_id).First(&sec)
	return sec
}

func (r MysqlStore) InsertSection(db *gorm.DB, config_id uint, section string) {
	cc := ConfigSection{ConfigID: config_id, Section: section}
	hh := r.GetSection(db, config_id, section)
	if hh.Section == "" {
		db.Create(&cc)
	}
}

func (r MysqlStore) UpdateSection(db *gorm.DB, sec ConfigSection) {
	db.Save(&sec)
}

func (r MysqlStore) DeleteSection(db *gorm.DB, sec ConfigSection) {
	db.Delete(&ConfigSection{}, sec.ID)
}

func (r MysqlStore) GetProperty(db *gorm.DB, config_id uint,
	sec_id uint, key string) ConfigProperty {
	var cc ConfigProperty
	db.Where("config_id = ? and  section_id = ? and s_key = ? ",
		config_id, sec_id, key).First(&cc)
	return cc
}

func (r MysqlStore) InsertProperty(db *gorm.DB, config_id uint, sec_id uint,
	key string, value string) {
	hh := r.GetProperty(db, config_id, sec_id, key)
	if hh.SKey == "" {
		pro := ConfigProperty{ConfigID: config_id,
			SectionID: sec_id, SKey: key, Value: value}
		db.Create(&pro)
	}

}
func (r MysqlStore) InsertPropertyObj(db *gorm.DB, pro ConfigProperty) {
	hh := r.GetProperty(db, pro.ConfigID, pro.SectionID, pro.SKey)
	if hh.SKey == "" {
		db.Create(&pro)
	}
}

func (r MysqlStore) UpdateProperty(db *gorm.DB, pro ConfigProperty) {
	db.Save(&pro)
}

func (r MysqlStore) DeleteProperty(db *gorm.DB, pro ConfigProperty) {
	db.Delete(&ConfigProperty{}, pro.ID)
}

func (r MysqlStore) GetPropertyValue(db *gorm.DB, cname string,
	section string, key string) ConfigProperty {
	conf := r.GetConfig(db, cname)
	if conf.CName == "" {
		return ConfigProperty{}
	}
	sec := r.GetSection(db, conf.ID, section)
	if sec.Section == "" {
		return ConfigProperty{}
	}
	return r.GetProperty(db, conf.ID, sec.ID, key)
}

/*
func (r MysqlStore) GetConfig(db *gorm.DB, cname string) {
	conf := r.GetConfig(db, cname)
	if conf.CName == "" {
		return ConfigProperty[]
	}
}
*/

func (r MysqlStore) InsertConfigValue(db *gorm.DB, cname string,
	section string, key string, value string) ConfigProperty {

	conf := r.GetConfig(db, cname)
	if conf.ID == 0 {
		r.InsertConfig(db, cname)
	}

	conf = r.GetConfig(db, cname)
	sec := r.GetSection(db, conf.ID, section)
	if sec.ID == 0 {
		r.InsertSection(db, conf.ID, section)
	}

	sec = r.GetSection(db, conf.ID, section)
	pro := r.GetProperty(db, conf.ID, sec.ID, key)
	if pro.ID == 0 {
		r.InsertProperty(db, conf.ID, sec.ID, key, value)
	} else {
		pro.Value = value
	}
	return pro
}

func (r MysqlStore) GetConfigValue(db *gorm.DB, cname string, section string,
	key string) ConfigProperty {

	conf := r.GetConfig(db, cname)
	if conf.ID == 0 {
		return ConfigProperty{}
	}

	sec := r.GetSection(db, conf.ID, section)
	if sec.ID == 0 {
		return ConfigProperty{}
	}

	pro := r.GetProperty(db, conf.ID, sec.ID, key)
	if pro.ID == 0 {
		return ConfigProperty{}
	}
	return pro
}

/*
config, section에 해당하는 모든 property들을 리턴한다.
*/
func (r MysqlStore) GetProperties(db *gorm.DB, cname string, section string) []ConfigProperty {
	ret := []ConfigProperty{}
	cc := r.GetConfig(db, cname)
	if cc.ID == 0 {
		return ret
	}
	ss := r.GetSection(db, cc.ID, section)
	if ss.ID == 0 {
		return ret
	}
	db.Where("config_id = ? and section_id = ? ", cc.ID, ss.ID).Find(&ret)
	return ret

}
