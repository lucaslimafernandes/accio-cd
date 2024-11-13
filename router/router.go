package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucaslimafernandes/accio-cd/handler"
	"github.com/lucaslimafernandes/accio-cd/utils"
)

func SetupRoutes(app *fiber.App, cdFile *utils.CDRunfile) {

	api := app.Group("")
	api.Get("/", handler.Hello)

	// api.Post("/webhook", handler.WH)
	api.Post("/webhook", func(c *fiber.Ctx) error {
		return handler.WH(c, cdFile)
	})
}
