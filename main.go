package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//Inicializa o Router utilizando as configs default do gin
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listens on 0.0.0.0:8080 by default
}
