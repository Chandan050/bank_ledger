package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"banking_ledger/internal/api/router"
	"banking_ledger/internal/models"
	"banking_ledger/internal/service"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	r := router.SetupRouter()

	account := models.Account{
		ID:      1,
		Balance: 100.0,
	}

	jsonValue, _ := json.Marshal(account)
	req, _ := http.NewRequest("POST", "/accounts", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAccount(t *testing.T) {
	r := router.SetupRouter()

	req, _ := http.NewRequest("GET", "/accounts/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateTransaction(t *testing.T) {
	r := router.SetupRouter()

	transaction := models.Transaction{
		AccountID:   1,
		Amount:      50,
		Type:        models.Debit,
		Description: "Test debit",
	}

	jsonValue, _ := json.Marshal(transaction)
	req, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetTransaction(t *testing.T) {
	r := router.SetupRouter()

	req, _ := http.NewRequest("GET", "/transactions/transaction_id_example", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetTransactionsInRange(t *testing.T) {
	r := router.SetupRouter()

	dateRange := models.Date{
		FromDate: "2023-01-01T00:00:00Z",
		ToDate:   "2023-12-31T23:59:59Z",
	}

	jsonValue, _ := json.Marshal(dateRange)
	req, _ := http.NewRequest("POST", "/transactions/range", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestProcessTransaction(t *testing.T) {
	ctx := context.Background()

	tx := models.Transaction{
		AccountID:   1,
		Amount:      50,
		Type:        models.Debit,
		Description: "Test debit processing",
	}

	_, err := service.ProcessTransaction(ctx, &tx)
	assert.NoError(t, err)
}
