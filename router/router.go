package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucaslimafernandes/accio-cd/handler"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")
	api.Get("/", handler.Hello)

}
