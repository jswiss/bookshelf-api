package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type createBookRequest struct {
	Title  string `json:"title" binding:required`
	Author string `json:"author" binding:required`
}

func (server *Server) createBook(ctx *gin.Context) {
	// var req createBookRequest
	// if err := ctx.ShouldBindJSON(&req); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	// 	return
	// }

	// arg := database.CreateBookParams{
	// 	Title:  req.Title,
	// 	Author: req.Author,
	// }

	// book, err := server.store.createBook(ctx, arg)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	// 	return
	// }

	// ctx.JSON(http.StatusOK, book)
}

func (server *Server) listBooks(ctx *gin.Context) {
	var req createBookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
}
