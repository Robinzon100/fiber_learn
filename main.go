package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	apiV1 "github.com/Robinzon100/fiber/api/v1"
)



type Post struct {
	gorm.Model
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}




var DB *gorm.DB
var posts = []Post{
	{Title: "something", Description: "something even more"},
	{Title: "something else", Description: "something else even more"},
}





func initDatabase() {
	dsn := "host=localhost user=admin password=admin dbname=oxeni port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	migrations(DB)
}


func migrations(DB *gorm.DB) {
	DB.AutoMigrate(&Post{})
	for i := range posts {
		DB.Create(&posts[i])
	}
}





func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
	})
	app.Use(cors.New())
	app.Use(logger.New())



	// =====> DB
	initDatabase()


	// =====> Api
	api := app.Group("/api")
	apiV1.Api_main(api)

	
	// =====> Server
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

func getPosts(c *fiber.Ctx) error {
	var posts []Post
	DB.Find(posts)
	return c.JSON(posts)
}
