package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	//Inicializa o Router utilizando as configs default do gin
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":6969")
}
