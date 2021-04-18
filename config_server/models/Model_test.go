package models_test

import (
	"fmt"
	"testing"

	"github.com/moonsungchul/configserver/models"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {

	store := models.MysqlStore{}
	db := store.Open("172.17.0.3", 3306, "fms_config", "moonstar", "wooag01")
	store.CreateTable(db)

	store.InsertConfig(db, "config_server")
	dat := store.GetConfig(db, "config_server")
	fmt.Println("config :", dat)
	assert.Equal(t, "config_server", dat.CName)

	store.InsertSection(db, dat.ID, "section1")
	ss := store.GetSection(db, dat.ID, "section1")
	fmt.Println("section :", ss)
	assert.Equal(t, "section1", ss.Section)

	store.InsertProperty(db, dat.ID, ss.ID, "key1", "value")
	pp := store.GetProperty(db, dat.ID, ss.ID, "key1")
	assert.Equal(t, "value", pp.Value)

	pro := store.GetPropertyValue(db, "config_server", "section1",
		"key1")
	fmt.Println(">>>>> pro : ", pro)
	assert.Equal(t, "value", pro.Value)

	store.UpdateConfig(db, dat)
	store.UpdateSection(db, ss)
	store.UpdateProperty(db, pp)

	store.DeleteConfig(db, dat)
	store.DeleteSection(db, ss)
	store.DeleteProperty(db, pp)
}

func Test2(t *testing.T) {

	store := models.MysqlStore{}
	db := store.Open("172.17.0.3", 3306, "fms_config", "moonstar", "wooag01")
	store.CreateTable(db)

	store.InsertConfigValue(db, "config1", "section1", "key1", "value1")
	val := store.GetConfigValue(db, "config1", "section1", "key1")
	fmt.Println("val : ", val)

	assert.Equal(t, "value1", val.Value)

	db.Exec("DELETE FROM configs")
	db.Exec("DELETE FROM config_sections")
	db.Exec("DELETE FROM config_properties")
}

func Test3(t *testing.T) {
	store := models.MysqlStore{}
	db := store.Open("172.17.0.3", 3306, "fms_config", "moonstar", "wooag01")
	store.CreateTable(db)

	store.InsertConfigValue(db, "config1", "section1", "key1", "value1")
	store.InsertConfigValue(db, "config1", "section1", "key2", "value2")
	store.InsertConfigValue(db, "config1", "section1", "key3", "value3")
	store.InsertConfigValue(db, "config1", "section1", "key4", "value4")

	store.InsertConfigValue(db, "config1", "section2", "key1", "value1")
	store.InsertConfigValue(db, "config1", "section2", "key2", "value2")
	store.InsertConfigValue(db, "config1", "section2", "key3", "value3")
	store.InsertConfigValue(db, "config1", "section2", "key4", "value4")

	ar := store.GetProperties(db, "config1", "section1")

	for i, obj := range ar {
		fmt.Println(i, obj)
		ss := fmt.Sprintf("value%d", i+1)
		assert.Equal(t, obj.Value, ss)
	}

	db.Exec("DELETE FROM configs")
	db.Exec("DELETE FROM config_sections")
	db.Exec("DELETE FROM config_properties")

}
