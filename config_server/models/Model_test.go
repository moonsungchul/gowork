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
