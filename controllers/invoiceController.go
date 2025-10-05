package controllers

import "github.com/gofiber/fiber/v3"

func GetInvoices(c *fiber.Ctx) {
	return c.SendString("get all invoice")
}

func GetInvoiceById(c *fiber.Ctx) {
	return c.SendString("get all invoice")
}

func CreateInvoice(c *fiber.Ctx) {
	return c.SendString("get all invoice")
}

func updateInvoice(c *fiber.Ctx) {
	return c.SendString("get all invoice")
}

func DeleteInvoice(c *fiber.Ctx) {
	return c.SendString("get all invoice")
}
