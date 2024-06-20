package main

import (
	"github.com/devesh/golang-react-jwt/database"
	"github.com/devesh/golang-react-jwt/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	config := cors.Config{
		AllowOrigins:     "http://localhost:5173/",
		AllowCredentials: true,
	}
	app.Use(cors.New(config))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello this server is open !")
	})

	routes.SetupRoutes(app)

	app.Listen(":8000")
}
