package routes

import (
	"github.com/gofiber/fiber/v2"
	"neeft_back/app/controllers/authController"
)

func SetupRouters(app *fiber.App) {

	//------------------ Auth ---------------------
	api := app.Group("/api")
	api.Post("login", authController.Login)

}
