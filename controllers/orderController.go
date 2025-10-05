package controllers

import "github.com/gofiber/fiber/v2"

func GetOrder(c *fiber.Ctx)error {
	return c.SendString("get all Order")
}

func GetOrderById(c *fiber.Ctx)error {
	return c.SendString("get all Order")
}

func AddOrder(c *fiber.Ctx)error {
	return c.SendString("get all Order")
}

func UpdateOrder(c *fiber.Ctx)error {
	return c.SendString("get all Order")
}

func DeleteOrder(c *fiber.Ctx) error{
	return c.SendString("get all Order")
}
