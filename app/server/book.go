package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	database "github.com/jswiss/bookshelf/app/database/sqlc"
	"github.com/lib/pq"
)

type createBookRequest struct {
	Title      string `json:"title" binding:"required"`
	Author     string `json:"author" binding:"required"`
	CoverImage string `json:"cover_image"`
}

type getBookRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type listBooksRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}

type updateBookRequest struct {
	ID         int32  `json:"id" binding:"required"`
	Title      string `json:"title" binding:"required"`
	Author     string `json:"author" binding:"required"`
	CoverImage string `json:"cover_image"`
}

func (server *Server) createBook(ctx *gin.Context) {
	var req createBookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := database.CreateBookParams{
		Title:      req.Title,
		Author:     req.Author,
		CoverImage: req.CoverImage,
	}

	book, err := server.store.CreateBook(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (server *Server) getBook(ctx *gin.Context) {
	var req getBookRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	book, err := server.store.GetBook(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (server *Server) listBooks(ctx *gin.Context) {
	var req listBooksRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := database.ListBooksParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	books, err := server.store.ListBooks(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (server *Server) updateBook(ctx *gin.Context) {
	var req updateBookRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := database.UpdateBookParams{
		ID:         req.ID,
		Title:      req.Title,
		Author:     req.Author,
		CoverImage: req.CoverImage,
	}

	err := server.store.UpdateBook(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	book, _ := server.store.GetBook(ctx, req.ID)

	ctx.JSON(http.StatusOK, book)
}
