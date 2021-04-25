package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/moonsungchul/configserver/models"
)

type RestConfig struct {
	Store *models.MysqlStore
}

type JConfig struct {
	ConfigName string `json:"config_name"`
	Section    string `json:"section"`
	Key        string `json:"key"`
	Value      string `json:"value"`
}

func (r RestConfig) AddConfig(c *gin.Context) {
	conf := JConfig{}
	err := c.BindJSON(&conf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(">>>> conf ", conf)
	db := r.Store.OpenDB()
	r.Store.InsertConfigValue(db, conf.ConfigName, conf.Section, conf.Key, conf.Value)
	c.JSON(200, gin.H{
		"msg": "add config value insert ok!",
	})

}
