package routers

import (
	"foodorder/controllers"

	"github.com/gofiber/fiber/v3"
)

func OrderRoutes(app *fiber.App) {
	order := app.Group("/order")

	order.Get("/", controllers.GetOrder)
	order.Get("/:id", controllers.GetOrderById)
	order.Post("/", controllers.AddOrder)
	order.Patch("/:id", controllers.UpdateOrder)
	order.Delete("/:id", controllers.DeleteOrder)

}
