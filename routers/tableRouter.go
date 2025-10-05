package routers

import (
	"foodorder/controllers"

	"github.com/gofiber/fiber/v3"
)

func TableRoutes(app *fiber.App) {
	table := app.Group("/table")

	table.Get("/", controllers.GetTable)
	table.Get("/:id", controllers.GetTableById)
	table.Post("/", controllers.AddTable)
	table.Patch("/:id", controllers.UpdateTable)
	table.Delete("/:id", controllers.DeleteTable)

}
