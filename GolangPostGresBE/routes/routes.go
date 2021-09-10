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
	app.Post("/api/getReimbursments", controllers.GetReimbursments)
	app.Post("/api/createReimbursment", controllers.CreateReimbursment)
	app.Post("/api/getHistory", controllers.GetHistory)
	app.Post("/api/getAllOpenReimbursments", controllers.GetAllOpenReimbursments)
	app.Post("/api/approveOrDeny", controllers.ApproveOrDeny)

}
