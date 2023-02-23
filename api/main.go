package main

import (
	"github.com/bthornhill123/go-auth-service/database"
	"github.com/bthornhill123/go-auth-service/handlers/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Post("/register", func(c *fiber.Ctx) error {
		return auth.Register(c)
	})
	app.Post("/login", func(c *fiber.Ctx) error {
		return auth.Login(c)
	})
	app.Post("/logout", func(c *fiber.Ctx) error {
		return auth.Logout(c)
	})
	app.Post("/authenticate", func(c *fiber.Ctx) error {
		return auth.Authenticate(c)
	})

	app.Listen(":8000")
}
