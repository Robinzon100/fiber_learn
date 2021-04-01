package apiV1

import (
	post "github.com/Robinzon100/fiber/api/v1/routes"
	"github.com/gofiber/fiber/v2"
) 

func Api_main(api fiber.Router) {
	apiv1 := api.Group("/v1")
	
	postsRoutes := apiv1.Group("/post")
	post.PostRoutes(postsRoutes)
}
