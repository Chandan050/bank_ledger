package service

import (
	"banking_ledger/internal/models"
	"banking_ledger/utils"
	"context"
)

func (h *Service) CreateAccount(ctx context.Context,
	account *models.Account) error {
	// Auto-generation of AccountId not implemented here due to missing DB access in service layer
	id := h.pg.GetMaxID(ctx)
	account.AccountId = utils.GenerateAccountNumber("6410", id)
	if err := h.pg.CreateAccount(ctx, account); err != nil {
		return err
	}
	return nil
}

func (h *Service) GetAccount(ctx context.Context,
	id string) (*models.Account, error) {
	account, err := h.pg.GetAccount(ctx, id)
	if err != nil {
		return nil, err
	}
	return account, nil
}
