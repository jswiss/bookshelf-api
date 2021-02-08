package api

import (
	"github.com/gin-gonic/gin"
)

// Server ...
type Server struct {
	router *gin.Engine
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// NewServer ...
func NewServer() *Server {
	server := &Server{}
	router := gin.Default()

	router.GET("/books", server.listBooks)
	router.POST("/books", server.createBook)

	server.router = router
	return server
}
