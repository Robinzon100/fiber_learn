package post

import "github.com/gofiber/fiber/v2"


func PostRoutes(postRoutes fiber.Router)  {
	postRoutes.Get("/all", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"someKey": "someValue",
		})
	})
}