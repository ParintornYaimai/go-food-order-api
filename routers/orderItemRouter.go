package routers

import (
	"foodorder/controllers"

	"github.com/gofiber/fiber/v2"
)

func OrderItemRoutes(app *fiber.App) {
	orderItem := app.Group("/order")

	orderItem.Get("/", controllers.GetOrderItem)
	orderItem.Get("/:id", controllers.GetOrderItemById)
	orderItem.Post("/", controllers.AddOrderItem)
	orderItem.Patch("/:id", controllers.UpdateOrderItem)
	orderItem.Delete("/:id", controllers.DeleteOrderItem)

}
