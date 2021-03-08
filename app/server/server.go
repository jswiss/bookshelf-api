package api

import (
	"github.com/gin-gonic/gin"
	database "github.com/jswiss/bookshelf/app/database/sqlc"
	"github.com/jswiss/bookshelf/util"
)

// Server ...
type Server struct {
	config util.Config
	store  database.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store database.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/books", server.createBook)
	router.GET("/books/:id", server.getBook)
	router.PUT("/books", server.updateBook)
	router.GET("/books", server.listBooks)
	router.DELETE("/books/:id", server.deleteBook)

	router.POST("/friends", server.createFriend)
	router.GET("/friends/:id", server.getFriend)
	router.PUT("/friends", server.updateFriend)
	router.GET("/friends", server.listFriends)
	router.DELETE("/friends/:id", server.deleteFriend)

	router.POST("/borrowed-books", server.createBorrowedBook)
	router.GET("/borrowed-books/:id", server.getBorrowedBook)
	router.PUT("/borrowed-books", server.updateBorrowedBook)
	router.GET("/borrowed-books", server.listBorrowedBooks)
	router.DELETE("/borrowed-books/:id", server.deleteBorrowedBook)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
