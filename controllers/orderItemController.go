package controllers

import "github.com/gofiber/fiber/v2"

func GetOrderItem(c *fiber.Ctx) error {
	return c.SendString("get all Order")
}

func GetOrderItemById(c *fiber.Ctx) error {
	return c.SendString("get all Order")
}

func AddOrderItem(c *fiber.Ctx) error {
	return c.SendString("get all Order")
}

func UpdateOrderItem(c *fiber.Ctx) error {
	return c.SendString("get all Order")
}

func DeleteOrderItem(c *fiber.Ctx) error {
	return c.SendString("get all Order")
}
