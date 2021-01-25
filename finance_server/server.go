package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/moonsungchul/finance/controllers"
)

func main() {
	gold := controllers.RestGold{}
	common := controllers.CommonInfo{}

	r := gin.Default()
	r.GET("/api/version", common.GetVersion)
	r.GET("/api/gold", gold.GetGoldData)
	fmt.Println("test")
	r.Run()
}
