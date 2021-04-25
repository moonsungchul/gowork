package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/moonsungchul/configserver/controllers"
	"github.com/moonsungchul/configserver/models"
)

func main() {
	store := models.MysqlStore{
		Host: "172.17.0.3", Port: 3306, Dbname: "fms_config",
		User: "moonstar", Passwd: "wooag01"}
	common := controllers.CommonInfo{}
	rconfig := controllers.RestConfig{&store}
	fmt.Println("start server ")

	r := gin.Default()
	r.GET("/api/version", common.GetVersion)
	r.POST("/api/config", rconfig.AddConfig)
	r.Run()

}
