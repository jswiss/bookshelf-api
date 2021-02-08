func TestCreateBook(t *testing.T) {
	arg := CreateBookParams{
		Title:      "Huckleberry Finn",
		Author:     "Mark Twain",
		CoverImage: "https://blogofgreengables.files.wordpress.com/2012/09/huckleberry-finn.jpg",
	}

	book, err := testQueries.CreateBook(context.Background(), arg)

}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, &sql.TxOptions{})
}
