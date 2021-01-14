package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jswiss/bookshelf/cmd/app/book"
	"gorm.io/gorm"
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
	database.dbConnection, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	app.Listen(3000)
}
