package controller

import (
	database "github.com/Robinzon100/fiber/database"
	"github.com/Robinzon100/fiber/database/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllPosts(c *fiber.Ctx) error {
	var posts []models.Post
	res := database.DBConn.Find(&posts)

	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   res.Error,
		})
	}

	return c.Status(fiber.StatusFound).JSON(fiber.Map{
		"allPosts": posts,
		"something": "something",
	})
}

// func PostPost(c *fiber.Ctx) error {
// 	db := database.DBConn

// 	post := new(models.Post)
// 	if err := c.BodyParser(post); err != nil {
// 		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
// 			"message":     "error",
// 			"description": "invalid data",
// 		})
// 	}

// 	// fmt.Println(post)

// 	db.Create(&post)

// 	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
// 		"message":    "success",
// 		"desciprion": "successfully added a post",
// 		"addedPost":  post,
// 	})
// }
