package api

import (
	db "BankstoreOpk/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Server struct {
	store *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// TODO: add routes to router
	router.POST("/accounts", server.CreateAccount)
	router.GET("/accounts/:id", server.GetAccount)
	router.GET("/accounts", server.ListAccounts)
	router.PUT("/update", server.UpdateAccount)
	router.DELETE("/delete/:id", server.DeleteAccount)
	server.router = router
	return server
}

// errorResponce return gin.H -> map[string]interface{}
func errorResponce(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// Start server method
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponce(err))
		return
	}
	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponce(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, account)
}