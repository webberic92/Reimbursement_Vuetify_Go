package routes

import (
	"example.com/m/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

}
