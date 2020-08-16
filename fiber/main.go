package main

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

func main() {
	app := fiber.New()
	app.Use(middleware.Recover())
	app.Use(middleware.Logger())

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	app.Get("/ping", func(c *fiber.Ctx) {
		c.JSON(fiber.Map{"message": "pong"})
	})

	log.Fatal(app.Listen(8080))
}
