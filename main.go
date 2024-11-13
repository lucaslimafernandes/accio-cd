package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lucaslimafernandes/accio-cd/router"
	"github.com/lucaslimafernandes/accio-cd/utils"
)

var CDFile *utils.CDRunfile

func init() {

	var err error

	CDFile, err = utils.CDReadRunfile("cdrunfile/on-push.yaml")
	if err != nil {
		log.Fatalln(err)
	}

}

func main() {

	fmt.Println("accio-cd")

	app := fiber.New()
	// app.Use(cors.New())

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))

}
