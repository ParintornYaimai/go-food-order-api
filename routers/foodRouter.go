package routers

import (
	"foodorder/controllers"

	"github.com/gofiber/fiber/v2"
)

func FoodRouter(app *fiber.App) {
	food := app.Group("/food")

	food.Get("/", controllers.GetFoods)
	food.Get("/:id", controllers.GetFoodById)
	food.Post("/", controllers.AddFood)
	food.Patch("/:id", controllers.UpdateFood)
}
