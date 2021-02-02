package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jswiss/bookshelf/app/dal"
	"github.com/jswiss/bookshelf/app/database"
	"github.com/jswiss/bookshelf/app/types"
)

// GetBooks returns the books list
func GetBooks(c *fiber.Ctx) error {
	d := &[]types.BookResponse{}

	err := dal.GetBooks(database.DB)
	if err != nil {
		return fiber.NewError(fiber.StatusConflict)
	}

	return c.JSON(&types.BooksResponse{
		Books: d,
	})
}
