package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	app := fiber.New()
	app.Use(middleware.Recover())
	app.Use(middleware.Logger())

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	app.Get("/ping", func(c *fiber.Ctx) {
		if err := c.JSON(fiber.Map{"message": "pong"}); err != nil {
			log.Println(err)
		}
	})

	log.Fatal(app.Listen(":" + port))
}
