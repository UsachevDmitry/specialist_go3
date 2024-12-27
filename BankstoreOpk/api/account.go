package api

import (
	db "BankstoreOpk/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/jackc/pgx/v5"
)

type CreateAccountRequest struct {
	Owner    string      `json:"owner" binding:"required"`
	Currency db.Currency `json:"currency" binding:"required,oneof=USR EUR"`
}

func (server *Server) CreateAccount(ctx *gin.Context) {
	var req CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponce(err))
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponce(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

type listAccountsRequest struct {
	//URL: /accounts/?page_id=3&page_size=10&page_id=3
	//Обработка query parameters в URL Path
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=15"`
}

func (server *Server) ListAccounts(ctx *gin.Context) {
	var req listAccountsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponce(err))
		return
	}
	arg := db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponce(err))
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}

// And PUT and DELETE urls and their endpoint

type UpdateAccountRequest struct {
	ID      int64 `json:"id" binding:"required"`
	Balance int64 `json:"balance"`
}

func (server *Server) UpdateAccount(ctx *gin.Context) {
	var req UpdateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponce(err))
	}

	arg := db.UpdateAccountParams{
		ID:      req.ID,
		Balance: req.Balance,
	}

	account, err := server.store.UpdateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponce(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

type DeleteAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) DeleteAccount(ctx *gin.Context) {
	var req DeleteAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponce(err))
		return
	}

	//id := req.ID

	rows, err := server.store.DeleteAccount(ctx, req.ID)
	if rows == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"Account not found": req.ID})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponce(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Deleted Account": req.ID})
}
