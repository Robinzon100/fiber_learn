package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/:id", func(c *fiber.Ctx) error {
		c.SendString(c.Params("id"))

		return c.JSON(fiber.Map{
			"name": c.Params("id"),
			"age":  20,
		})
	})

	app.Listen(":3000")
}
