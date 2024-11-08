package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lucaslimafernandes/accio-cd/router"
)

func main() {

	fmt.Println("accio-cd")

	app := fiber.New()
	// app.Use(cors.New())

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))

}
