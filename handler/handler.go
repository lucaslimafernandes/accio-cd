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

	eventType := c.Get("X-GitHub-Event", "None")

	// Parse o corpo JSON do webhook
	err := c.BodyParser(&payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Erro ao processar o webhook")
	}

	// Exiba o payload do webhook
	fmt.Printf("Recebido webhook: %v\n", payload)
	fmt.Printf("\n\n%v", string(c.Body()))

	fmt.Printf("\n\n%v", eventType)

	return c.SendString("Webhook recebido com sucesso!")
}
