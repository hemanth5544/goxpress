package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/hemanth5544/goxpress/initializers"
	"github.com/hemanth5544/goxpress/internal/db"
)

func main() {
	app := fiber.New()
	initializers.LoadEnv()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg": "Hello Hemanth this side",
		})
	})
	db.ConnectDatabase()

	fmt.Println("new server is created")
	app.Listen(":3000")
}
