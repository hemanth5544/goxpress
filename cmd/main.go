package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hemanth5544/goxpress/initializers"
	"github.com/hemanth5544/goxpress/internal/db"
)

func main() {
	initializers.LoadEnv()
	db.ConnectDatabase()

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {

		//this "c" will have this /ping route all req,res c is a context

		//c.ShouldBindJSON  will map the body paylod wiht the ree and contevert them to sturct ad the  incoing req was json'
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	port := os.Getenv("AP_PORT")

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}

	fmt.Println("new server is created")
}
