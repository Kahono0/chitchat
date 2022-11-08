package main

import (
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/kahono922/chitchat/database"
	"github.com/kahono922/chitchat/router"
	"github.com/joho/godotenv"
)

func loadEnv(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	app := fiber.New()
	loadEnv()
	port := os.Getenv("PORT")

	//connect to database
	database.Conn()

	//set up routes
	router.SetUpRoutes(app)

	//start server
	log.Println("Server is running on port :", port)
	log.Fatal(app.Listen(":"+port))

}
