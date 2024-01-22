package main

import (
	"github.com/gin-gonic/gin"
)
func main() {
	router := gin.Default() 

	router.GET("/ping", func(c *gin.Context) { 
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")  // Starts the Gin server on the default port (8080)
}