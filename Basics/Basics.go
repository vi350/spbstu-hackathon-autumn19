package Basics

import (
	"github.com/gin-gonic/gin"
	"github.com/vi350/spbstu-hackathon-autumn19/DB"
)

func Welcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"API работает":              true,
		"Успешное подключение к бд": DB.Status(),
	})
}
