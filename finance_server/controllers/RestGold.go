package controllers

import "github.com/gin-gonic/gin"

type RestGold struct {
}

func (r RestGold) GetGoldData(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "wo hhhaaaa",
	})
}
