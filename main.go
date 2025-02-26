package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/persons", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"persons": []fiber.Map{
				{
					"id":   1,
					"name": "John Doe",
				},
				{
					"id":   2,
					"name": "Jane Doe",
				},
				{
					"id":   3,
					"name": "John Smith",
				},
			},
		})
	})

	app.Post("/persons", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Person created successfully",
		})
	})

	log.Info("Starting server on port 8080")
	app.Listen(":8080")
}
