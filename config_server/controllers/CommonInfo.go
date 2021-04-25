package controllers

import "github.com/gin-gonic/gin"

type CommonInfo struct {
}

func (r CommonInfo) GetVersion(c *gin.Context) {
	c.JSON(200, gin.H{
		"version": "finance server api 0.1",
	})
}
