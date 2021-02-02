package dal

import (
	"database/sql"
	"fmt"

	"github.com/jswiss/bookshelf/app/types"
)

// GetBooks ...
func GetBooks(db *sql.DB) []types.BookResponse {
	rows, err := db.Query("SELECT * FROM books")
	defer rows.Close()
	books := []types.BookResponse{}
	for rows.Next() {
		var b types.BookResponse
		books = append(books, b)
	}
	if err != nil {
		fmt.Printf("There is an error")
	}
	return books
}
