package routes

import (
	"github.com/devesh/golang-react-jwt/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api") // grouping the api route

	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)
	api.Get("/getuser", controllers.GetUser)
	api.Post("/logout", controllers.Logout)
}
