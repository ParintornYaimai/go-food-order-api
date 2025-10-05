package controllers

import "github.com/gofiber/fiber/v2"

func GetMenu(c *fiber.Ctx) error{
	return c.SendString("get all menu")
}

func GetMenuById(c *fiber.Ctx)error {
	return c.SendString("get all menu")
}

func AddMenu(c *fiber.Ctx) error{
	return c.SendString("get all menu")
}

func UpdateMenu(c *fiber.Ctx) error{
	return c.SendString("get all menu")
}

func DeleteMenu(c *fiber.Ctx)error {
	return c.SendString("get all menu")
}
