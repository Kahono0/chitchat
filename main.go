package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kahono922/chitchat/database"
	"github.com/kahono922/chitchat/router"
)

func main() {
	app := fiber.New()

	//connect to database
	database.Conn()

	//set up routes
	router.SetUpRoutes(app)

	//start server
	log.Fatal(app.Listen(":3000"))

}
