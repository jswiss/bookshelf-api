package database

import (
	"context"
	"testing"

	"github.com/jswiss/bookshelf/util"
	"github.com/stretchr/testify/require"
)

func createRandomBook(t *testing.T) Book {
	arg := CreateBookParams{
		Title:  util.RandomTitle(),
		Author: util.RandomAuthor()}

	book, err := testQueries.CreateBook(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, book)

	require.Equal(t, arg.Title, book.Title)
	require.Equal(t, arg.Author, book.Author)

	require.NotZero(t, book.ID)
	require.NotZero(t, book.CreatedAt)

	return book
}

func TestCreateBook(t *testing.T) {
	createRandomBook(t)
}
