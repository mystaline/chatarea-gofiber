package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func setupTestApp() *fiber.App {
	app := fiber.New()
	app.Get("/ping", GetUsers)
	app.Post("/api/v1/auth/register", Register)
	app.Post("/api/v1/auth/login", Login)
	app.Get("/api/v1/me/profile", Me)
	app.Put("/api/v1/me/profile", EditProfile)
	return app
}
