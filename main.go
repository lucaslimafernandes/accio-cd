package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lucaslimafernandes/accio-cd/logs"
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

	logs.InitDB()
	// defer logs.DB.Close()

}

func main() {

	fmt.Println("accio-cd")

	app := fiber.New()
	// app.Use(cors.New())

	router.SetupRoutes(app, CDFile)
	log.Fatal(app.Listen(":3000"))

}
