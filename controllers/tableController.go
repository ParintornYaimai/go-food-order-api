package controllers

import "github.com/gofiber/fiber/v2"

func GetTable(c *fiber.Ctx) error {
	c.Status(200)                    // ตั้ง status
	return c.SendString("user data") // คืนค่า error ตาม signature
}

func GetTableById(c *fiber.Ctx) error {
	return c.SendString("get all Table")
}

func AddTable(c *fiber.Ctx) error {
	return c.SendString("get all Table")
}

func UpdateTable(c *fiber.Ctx) error {
	return c.SendString("get all Table")
}

func DeleteTable(c *fiber.Ctx) error {
	return c.SendString("get all Table")
}
