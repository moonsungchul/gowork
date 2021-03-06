package controllers

import (
	"fmt"
	"net/http"

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

func (r RestConfig) GetConfig(c *gin.Context) {
	cname := c.Params.ByName("cname")
	section := c.Params.ByName("section")
	key := c.Params.ByName("key")

	db := r.Store.OpenDB()
	pro := r.Store.GetConfigValue(db, cname, section, key)
	c.JSON(http.StatusOK, pro)

}

func (r RestConfig) DeleteProperty(c *gin.Context) {
	cname := c.Params.ByName("cname")
	section := c.Params.ByName("section")
	key := c.Params.ByName("key")

	db := r.Store.OpenDB()
	pro := r.Store.GetConfigValue(db, cname, section, key)
	if pro.ID == 0 {
		msg := Msg{MsgID: -100, Message: "지정된 프로퍼티를 찾을 수 없습니다. "}
		c.JSON(http.StatusOK, msg)
	}
	msg := Msg{MsgID: 100, Message: "지정된 프로퍼티를 삭제 했습니다. "}
	fmt.Println("@@@@@@@@@@ : ", msg)
	c.JSON(http.StatusOK, msg)
}
