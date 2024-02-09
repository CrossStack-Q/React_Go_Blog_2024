package main

import (
	"log"
	"os"

	"github.com/CrossStack-Q/React_Go_Blog_2024/backend/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	app.Listen(":" + port)
}
