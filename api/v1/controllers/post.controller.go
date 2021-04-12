package post

import (
	database "github.com/Robinzon100/fiber/database"
	table "github.com/Robinzon100/fiber/database/migrations/tables"
	"github.com/gofiber/fiber/v2"
)




func GetPosts(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"someKey": "someValue",
	})
}




func PostPost(c *fiber.Ctx) error {
	db := database.DBConn

	post := new(table.Post)
	if err := c.BodyParser(post); err != nil {
		c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"message":     "error",
			"description": "invalid data",
		})
	}

	postedPost := db.Create(post)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":    "success",
		"desciprion": "successfully added a post",
		"addedPost":  postedPost,
	})
}
