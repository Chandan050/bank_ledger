package router

import (
	"banking_ledger/internal/api/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, h *handler.Handler) *gin.Engine {

	router.GET("/account/:id", h.GetAccount)
	router.POST("/account", h.CreateAccount)
	router.POST("/transaction", h.CreateTransaction)
	router.GET("/transactions/:transaction_id", h.GetTransaction)
	router.GET("/transactions_list/:account_id", h.GetTransactionsInRange)
	return router

}
