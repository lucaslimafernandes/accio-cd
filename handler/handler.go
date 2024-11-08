package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}

func WH(c *fiber.Ctx) error {
	var payload map[string]interface{}

	// Parse o corpo JSON do webhook
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Erro ao processar o webhook")
	}

	// Exiba o payload do webhook
	fmt.Printf("Recebido webhook: %v\n", payload)

	return c.SendString("Webhook recebido com sucesso!")
}
