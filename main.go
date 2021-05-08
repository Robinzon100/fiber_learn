package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	apiV1 "github.com/Robinzon100/fiber/api/v1"
	dbUtils "github.com/Robinzon100/fiber/database/utils"
)







func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
	})
	app.Use(cors.New())
	app.Use(logger.New())


	// =====> DB
	dbUtils.InitDatabase()
	// defer database.DBConn.()
	// defer database.DB.Close()


	// =====> Api
	api := app.Group("/api")
	apiV1.Api_main(api)

	// =====> Server
	err := app.Listen(":5000")
	if err != nil {
		panic(err)
	}
}
