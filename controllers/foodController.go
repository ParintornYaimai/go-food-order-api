package controllers

import "github.com/gofiber/fiber/v3"

func GetFoods(c *fiber.Ctx) error {
	return c.SendString("🍜 All foods list")
}

func GetFoodById(c *fiber.Ctx) error {
	return c.SendString("🍜 All foods list")
}

func AddFood(c *fiber.Ctx) error {
	return c.SendString("🍜 All foods list")
}

func UpdateFood(c *fiber.Ctx) error {
	return c.SendString("🍜 All foods list")
}
