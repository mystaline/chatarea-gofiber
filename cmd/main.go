package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mystaline/chatarea-gofiber/internal/config"
	route "github.com/mystaline/chatarea-gofiber/internal/router"
)

func main() {
	app := fiber.New()

	// Optional: Connect DB
	config.ConnectDB()

	// Register routes
	route.SetupRoutes(app)

	fmt.Println(os.Getenv("APP_PORT"))
	port := os.Getenv("APP_PORT")
	log.Println("Server running at http://localhost:" + port)
	log.Fatal(app.Listen(":" + port))
}
