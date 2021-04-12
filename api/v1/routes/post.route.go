package post

import (
	postController "github.com/Robinzon100/fiber/api/v1/controllers"
	"github.com/gofiber/fiber/v2"
) 


func PostRoutes(postRoutes fiber.Router)  {
	postRoutes.Get("/all", postController.GetPosts)
	postRoutes.Post("/post", postController.PostPost)
}