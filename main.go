package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/lucaslimafernandes/accio-cd/router"
)

func main() {

	fmt.Println("accio-cd")

	app := fiber.New()
	app.Use(cors.New())

	router.SetupRoutes(app)

}
