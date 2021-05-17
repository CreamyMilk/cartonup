package main

import (
	"fmt"

	"github.com/CreamyMilk/cartonup/database"
	"github.com/CreamyMilk/cartonup/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: false,
	})
	app.Use(cors.New())
	router.SetupRoutes(app)
	if err := database.Connect(); err != nil {
		fmt.Printf("DB ERROR %v", err)
	}
	defer database.DB.Close()
	app.Listen(":1337")
}
