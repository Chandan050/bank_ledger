package service

import (
	"banking_ledger/internal/models"
	"banking_ledger/internal/repository"
	"banking_ledger/internal/repository/postgres"
	"banking_ledger/utils"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Service struct {
	pg    repository.AccountRepository
	mongo repository.TxRepository
}

func NewService(pg repository.AccountRepository, mongo repository.TxRepository) *Service {
	return &Service{
		pg:    pg,
		mongo: mongo,
	}
}

func (s *Service) CreateTransaction(ctx context.Context, tx *models.Transaction) (any, error) {
	// Validate transaction type and amount
	if tx.Amount <= 0 {
		return nil, errors.New("transaction amount must be positive")
	}
	if tx.Type != models.Debit && tx.Type != models.Credit {
		return nil, errors.New("invalid transaction type")
	}

	return s.ProcessTransaction(ctx, tx)
}

func (s *Service) ProcessTransaction(con_text context.Context, tx *models.Transaction) (any, error) {
	ctx, cancel := context.WithTimeout(con_text, 3*time.Second)
	defer cancel()

	err := s.pg.(*postgres.PostgresAccountRepository).Db.Transaction(func(txDB *gorm.DB) error {
		account, err := s.pg.GetAccount(ctx, tx.AccountID)
		if err != nil {
			return err
		}

		switch tx.Type {
		case models.Debit:
			if account.Balance < float32(tx.Amount) {
				return errors.New("insufficient balance")
			}
			account.Balance -= float32(tx.Amount)
		case models.Credit:
			account.Balance += float32(tx.Amount)
		default:
			return errors.New("invalid transaction type")
		}

		// Update account balance atomically
		err = s.pg.UpdateAccountBalance(ctx, account.ID, account.Balance)
		if err != nil {
			return err
		}

		// Log transaction in MongoDB
		_, err = s.mongo.CreateTransaction(ctx, tx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *Service) GetTransaction(ctx context.Context, id string) (*models.Transaction, error) {
	return s.mongo.GetTansaction(ctx, id)
}

func (s *Service) GetTransactions(ctx context.Context, pd *models.Date,
	account_id string) (*[]models.Transaction, error) {
	if pd == nil {
		return nil, errors.New("error of dates are empty")
	}

	fromDate, err := utils.ParseTime(pd.FromDate)
	if err != nil {
		return nil, errors.New("error invalide fromdate")
	}

	toDate, err := utils.ParseTime(pd.ToDate)
	if err != nil {
		return nil, errors.New("error, invalide todate")
	}
	return s.mongo.GetTransactions(ctx, fromDate, toDate, account_id)
}
