package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hemanth5544/goxpress/initializers"
	"github.com/hemanth5544/goxpress/internal/db"
)

func main() {
	initializers.LoadEnv()
	db.ConnectDatabase()

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() 

	fmt.Println("new server is created")
}
