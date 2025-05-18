package handler

import (
	"banking_ledger/internal/models"
	"banking_ledger/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{
		svc: *svc,
	}
}

func (h *Handler) CreateAccount(c *gin.Context) {
	var account models.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"message": "incorrect data", "err": err.Error()})
		return
	}

	err := h.svc.CreateAccount(c.Request.Context(), &account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)

}

func (h *Handler) GetAccount(c *gin.Context) {
	idStr := c.Param("id")
	account, err := h.svc.GetAccount(c.Request.Context(), idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)

}
