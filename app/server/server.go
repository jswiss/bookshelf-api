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
	router.GET("/books/book/:id", server.getBook)
	router.GET("/books", server.listBooks)

	// router.POST("/friends", server.CreateFriend)
	// router.GET("/friends/:id", server.getFriend)
	// router.GET("/friends", server.getFriends)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
