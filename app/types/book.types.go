package types

// BookResponse struct contains the book field which should be returned in a response
type BookResponse struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	CoverImage string `json:"cover_image"`
	InStock    bool   `json:"in_stock"`
}

// BookCreateResponse struct defines the /book/create response
type BookCreateResponse struct {
	Book *BookResponse `json:"book"`
}

// BooksResponse defines the books list
type BooksResponse struct {
	Books *[]BookResponse `json:"books"`
}
