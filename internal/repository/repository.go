package repository

import (
	"banking_ledger/internal/models"
	"context"
	"time"
)

type TxRepository interface {
	CreateTransaction(ctx context.Context, tx *models.Transaction) (any, error)
	GetTansaction(ctx context.Context, id string) (*models.Transaction, error)
	GetTransactions(ctx context.Context, fromDate, toDate *time.Time, account_id string) (*[]models.Transaction, error)
}

type AccountRepository interface {
	CreateAccount(ctx context.Context, acc *models.Account) error
	GetAccount(ctx context.Context, id string) (*models.Account, error)
	UpdateAccountBalance(ctx context.Context, accountID uint, newBalance float32) error
	GetMaxID(ctx context.Context) (maxID int)
}
