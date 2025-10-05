package controllers

import "github.com/gofiber/fiber/v3"

func GetMenu(c *fiber.Ctx) {
	return c.SendString("get all menu")
}

func GetMenuById(c *fiber.Ctx) {
	return c.SendString("get all menu")
}

func AddMenu(c *fiber.Ctx) {
	return c.SendString("get all menu")
}

func UpdateMenu(c *fiber.Ctx) {
	return c.SendString("get all menu")
}

func DeleteMenu(c *fiber.Ctx) {
	return c.SendString("get all menu")
}
