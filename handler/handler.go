package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/lucaslimafernandes/accio-cd/utils"
)

func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}

func WH(c *fiber.Ctx, cdFile *utils.CDRunfile) error {

	// var payload map[string]interface{}
	var payload WebhookPayload

	eventType := c.Get("X-GitHub-Event", "None")

	// Parse o corpo JSON do webhook
	err := c.BodyParser(&payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Erro ao processar o webhook")
	}

	// Exiba o payload do webhook
	// fmt.Printf("Recebido webhook: %v\n", payload)
	// fmt.Printf("\n\n%v", string(c.Body()))

	fmt.Printf("\n\neventType\t%v\n\n", eventType)

	if verifyURL(cdFile, &payload) {

		msg := verifyEvent(&eventType, cdFile)

		response := fmt.Sprintf("Webhook recebido com sucesso!\t%v", msg)
		return c.SendString(response)

	}

	return c.SendString("Webhook recebido com sucesso!    - Sem detalhes")

}

func verifyURL(cdFile *utils.CDRunfile, payload *WebhookPayload) bool {
	return cdFile.GitUrl == payload.Repository.SSHURL
}

func verifyEvent(event *string, cdFile *utils.CDRunfile) string {

	var message string
	if *event == cdFile.On {

		switch cdFile.On {
		case "push":
			message = "EventPush()"
			EventPush(cdFile)
		default:
			message = "DEFAULT"
			fmt.Println("DEFAULT")
		}

		fmt.Printf("\nAQUI\n")
		fmt.Printf("%v - %v\n", *event, cdFile.On)
	} else {
		fmt.Printf("\n%v - %v\n", *event, cdFile.On)
	}

	return message

}
