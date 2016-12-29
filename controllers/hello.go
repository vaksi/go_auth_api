/*
=====================Golang======================
	Author 	: Audy Vaksi Pranata
	Link	: https://github.com/vaksi
	Email	: vaksipranata@gmail.com
==================================================
*/
package controllers

import (
	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"text": "Hello World.",
	})
}
