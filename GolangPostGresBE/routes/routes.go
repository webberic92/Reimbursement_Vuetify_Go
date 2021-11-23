package routes

import (
	"example.com/m/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Put("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Get("/api/logout", controllers.Logout)
	app.Get("/api/getReimbursments", controllers.GetReimbursments)
	app.Post("/api/createReimbursment", controllers.CreateReimbursment)
	app.Get("/api/getHistory", controllers.GetHistory)
	app.Get("/api/getAllOpenReimbursments", controllers.GetAllOpenReimbursments)
	app.Put("/api/approveOrDeny", controllers.ApproveOrDeny)
	app.Get("/api/getAllHistory", controllers.GetAllHistory)

}
