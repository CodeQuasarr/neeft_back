package routes

import (
	"github.com/gofiber/fiber/v2"
	"neeft_back/app/controllers/authController"
	"neeft_back/app/controllers/users"
)

func SetupRouters(app *fiber.App) {

	//------------------ Auth ---------------------
	api := app.Group("/api")
	api.Post("login", authController.Login)

	//------------------ Users ------------------
	api.Post("/users", users.CreateUser)
	api.Get("/users", users.GetAllUser)
	api.Get("/user/:id", users.GetUser)
	api.Put("/user/:id", users.UpdateUser)
	api.Delete("/user/:id", users.DeleteUser)

}
