package controllers

import (
	"github.com/gin-gonic/gin"
)

func CheckHandler(c *gin.Context) {

	c.JSON(200, gin.H{
		"Allow": "true",
	})
}
