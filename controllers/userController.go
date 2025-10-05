package controllers

import "github.com/gofiber/fiber/v2"

func GetUserDataById(c *fiber.Ctx) error {
	return c.SendString("user data")
}

func Register(c *fiber.Ctx) error {
	return c.SendString("user data")
}

func Login(c *fiber.Ctx) error {
	return c.SendString("user data")
}

func Update(c *fiber.Ctx) error {
	return c.SendString("user data")
}

func DeleteAccount(c *fiber.Ctx) error {
	return c.SendString("user data")
}
