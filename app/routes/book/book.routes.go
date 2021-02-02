package routes

import (
	"github.com/gofiber/fiber"
	"github.com/jswiss/bookshelf/app/services"
)

// BookRoutes ...
func BookRoutes(app fiber.Router) {
	r := app.Group("/books")
	r.Get("/list", services.GetBooks)
}
