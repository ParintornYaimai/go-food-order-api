package routers

import (
	"foodorder/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	user := app.Group("/user")

	user.Get("/:id", controllers.GetUserDataById)
	user.Post("/register", controllers.Register)
	user.Post("/login", controllers.Login)
	user.Patch("/update", controllers.Update)
	user.Delete("/deleteAccount", controllers.DeleteAccount)
}
