package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jswiss/bookshelf/database"
	"github.com/jswiss/bookshelf/routes/book"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DbConnection()
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
}

func main() {
	initDatabase()
	app := fiber.New()

	setupRoutes(app)
	app.Listen(3000)
}
