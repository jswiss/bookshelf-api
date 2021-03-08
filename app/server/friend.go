package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	database "github.com/jswiss/bookshelf/app/database/sqlc"
	"github.com/lib/pq"
)

type createFriendRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Photo    string `json:"photo" binding:"required"`
}

type getFriendRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type listFriendsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}

type updateFriendRequest struct {
	ID       int32  `json:"id" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	Photo    string `json:"photo" binding:"required"`
}

func (server *Server) createFriend(ctx *gin.Context) {
	var req createFriendRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := database.CreateFriendParams{
		FullName: req.FullName,
		Photo:    req.Photo,
	}

	book, err := server.store.CreateFriend(ctx, arg)
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

func (server *Server) getFriend(ctx *gin.Context) {
	var req getFriendRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	book, err := server.store.GetFriend(ctx, req.ID)
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

func (server *Server) listFriends(ctx *gin.Context) {
	var req listFriendsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := database.ListFriendsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	books, err := server.store.ListFriends(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (server *Server) updateFriend(ctx *gin.Context) {
	var req updateFriendRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := database.UpdateFriendParams{
		ID:       req.ID,
		FullName: req.FullName,
		Photo:    req.Photo,
	}

	err := server.store.UpdateFriend(ctx, arg)
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

	book, _ := server.store.GetFriend(ctx, req.ID)

	ctx.JSON(http.StatusOK, book)
}

func (server *Server) deleteFriend(ctx *gin.Context) {
	var req getFriendRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteFriend(ctx, req.ID)
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
