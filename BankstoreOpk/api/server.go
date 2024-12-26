package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	db "BankstoreOpk/db/sqlc"
)

type Server struct {
	store *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	//TODO: add routes to router
	router.POST("/accounts", server.CreateAccount)

	server.router = router
	return server

}

type CreateAccountReques struct {
	Owner string `json:"owner" binding:"required"`
	Currence db.Currency `json:"currency" binding:"required,oneof=USR EUR"`
}

func (server *Server) CreateAccount(ctx *gin.Context) {
	//TODO
}

// errorResponse return gin.H -> map[string]interface()
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

//Start server func
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
