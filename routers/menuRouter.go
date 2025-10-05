package routers

import (
	"foodorder/controllers"

	"github.com/gofiber/fiber/v3"
)

func MenuRouters(app *fiber.App) {
	menu := app.Group("/menu")

	menu.Get("/", controllers.GetMenu)
	menu.Get("/:id", controllers.GetMenuById)
	menu.Post("/", controllers.AddMenu)
	menu.Patch("/:id", controllers.UpdateMenu)
	menu.Delete("/:id", controllers.DeleteMenu)
}
