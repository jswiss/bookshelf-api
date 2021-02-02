package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"

	"github.com/jswiss/bookshelf/app/database"
	routes "github.com/jswiss/bookshelf/app/routes/book"
)

func helloWorld(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"msg": "Hello, World!"})
}

func initDatabase() {
	var err error
	database.Connect()
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
}

func main() {
	initDatabase()
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	routes.BookRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		if c.Is("json") {
			return c.Next()
		}
		return c.SendString("Only JSON allowed!")
	})
	app.Listen("3000")
}
