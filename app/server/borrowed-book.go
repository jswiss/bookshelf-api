package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	database "github.com/jswiss/bookshelf/app/database/sqlc"
	"github.com/lib/pq"
)

type createBorrowedBookRequest struct {
	BookID   int32 `json:"book_id" binding:"required"`
	FriendID int32 `json:"friend_id" binding:"required"`
}

type getBorrowedBookRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type listBorrowedBooksRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}

type updateBorrowedBookRequest struct {
	ID int32 `json:"id" binding:"required"`
}

func (server *Server) createBorrowedBook(ctx *gin.Context) {
	var req createBorrowedBookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := database.CreateBorrowedBookParams{
		BookID:   req.BookID,
		FriendID: req.FriendID,
	}

	book, err := server.store.CreateBorrowedBook(ctx, arg)
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

func (server *Server) getBorrowedBook(ctx *gin.Context) {
	var req getBorrowedBookRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	book, err := server.store.GetBorrowedBook(ctx, req.ID)
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

func (server *Server) listBorrowedBooks(ctx *gin.Context) {
	var req listBorrowedBooksRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := database.ListBorrowedBooksParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	books, err := server.store.ListBorrowedBooks(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (server *Server) updateBorrowedBook(ctx *gin.Context) {
	var req updateBorrowedBookRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.UpdateBorrowedBook(ctx, req.ID)
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

	borrowedBook, _ := server.store.GetBorrowedBook(ctx, req.ID)

	ctx.JSON(http.StatusOK, borrowedBook)
}

func (server *Server) deleteBorrowedBook(ctx *gin.Context) {
	var req getBorrowedBookRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteBorrowedBook(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, req.ID)
}
