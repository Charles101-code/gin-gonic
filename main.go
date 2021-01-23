package main

import (
	"gin-gonic/config"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Init configuration to load conf data
	config.Init()

	serverConf := config.ConfType.Server
	port := serverConf.Port
	r.Run(port) // listen and serve on 0.0.0.0:8081
}
