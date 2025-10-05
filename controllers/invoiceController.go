package controllers

import "github.com/gofiber/fiber/v2"

func GetInvoices(c *fiber.Ctx) error {
	return c.SendString("get all invoice")
}

func GetInvoiceById(c *fiber.Ctx) error {
	return c.SendString("get all invoice")
}

func CreateInvoice(c *fiber.Ctx) error {
	return c.SendString("get all invoice")
}

func UpdateInvoice(c *fiber.Ctx) error {
	return c.SendString("get all invoice")
}

func DeleteInvoice(c *fiber.Ctx) error {
	return c.SendString("get all invoice")
}
