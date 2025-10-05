package main

import (
	"log"
	"os"

	"foodorder/database"
	"foodorder/middleware"
	"foodorder/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Connect to the database
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected successfully")

	// middleware
	app.Use(logger.New())
	app.Use(middleware.AuthMiddleware)

	//routes
	routers.FoodRouter(app)
	routers.InvoiceRouter(app)
	routers.MenuRouters(app)
	routers.OrderItemRoutes(app)
	routers.OrderRoutes(app)
	routers.TableRoutes(app)
	routers.UserRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
