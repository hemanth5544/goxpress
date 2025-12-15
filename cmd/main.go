package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hemanth5544/goxpress/initializers"
	"github.com/hemanth5544/goxpress/internal/auth"
	"github.com/hemanth5544/goxpress/internal/cart"
	"github.com/hemanth5544/goxpress/internal/db"
	"github.com/hemanth5544/goxpress/internal/order"
	"github.com/hemanth5544/goxpress/internal/product"
)

func main() {
	initializers.LoadEnv()
	db, err := db.ConnectDatabase()
	if err != nil {
		log.Println("Failed to connect database")
	}
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {

	//?this "c" will have this /ping route all req,res c is a context

	//?c.ShouldBindJSON  will map the body paylod wiht the ree and contevert them to sturct ad the  incoing req was json'
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//loading our routes
	auth.SetupAuth(router, db)
	product.ProductRouter(router, db)
	cart.SetupCart(router, db)
	order.SetupOrder(router, db)

	port := os.Getenv("APP_PORT")

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}

	fmt.Println("new server is created")
}
