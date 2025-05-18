package handler

import (
	"banking_ledger/internal/models"
	"banking_ledger/queue"
	"encoding/json"
	"log"
	"net/http"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	marshalleddata, err := json.Marshal(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal transaction"})
		return
	}
	message := &sarama.ProducerMessage{
		Topic: "createTransaction",
		Value: sarama.ByteEncoder(marshalleddata),
	}

	partiation, offset, err := queue.StartProducer().SendMessage(message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message to queue"})
		return

	}
	log.Printf("Transaction created with partition %d and offset %d", partiation, offset)

	c.JSON(http.StatusAccepted, gin.H{"message": "Transaction created"})

}

func (h *Handler) GetTransactionsInRange(c *gin.Context) {
	var dateRange models.Date
	account_id := c.Param("account_id")
	if err := c.ShouldBindJSON(&dateRange); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	transactions, err := h.svc.GetTransactions(ctx, &dateRange, account_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (h *Handler) GetTransaction(c *gin.Context) {
	transactionID := c.Param("transaction_id")
	ctx := c.Request.Context()

	transaction, err := h.svc.GetTransaction(ctx, transactionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaction)
}
