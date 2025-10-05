package routers

import (
	"foodorder/controllers"

	"github.com/gofiber/fiber/v2"
)

func InvoiceRouter(app *fiber.App) {
	invoice := app.Group("/invoice")

	invoice.Get("/", controllers.GetInvoices)
	invoice.Get("/:id", controllers.GetInvoiceById)
	invoice.Post("/", controllers.CreateInvoice)
	invoice.Patch("/:id", controllers.UpdateInvoice)
	invoice.Delete("/:id", controllers.DeleteInvoice)

}
