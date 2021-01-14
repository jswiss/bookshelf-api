package book

import (
	"github.com/gofiber/fiber"
)

func GetBooks(c *fiber.Ctx) {
	c.Send("All Books")
}

func GetBook(c *fiber.Ctx) {
	c.Send("Single Book")
}

func NewBook(c *fiber.Ctx) {
	c.Send("New Book")
}

func DeleteBook(c *fiber.Ctx) {
	c.Send("Delete Book")
}
