package postgres

import (
	"banking_ledger/internal/models"
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresAccountRepository struct {
	Db *gorm.DB
}

func NewPostgresAccountRepository() (*PostgresAccountRepository, error) {
	user_name := os.Getenv("user_name")
	user_password := os.Getenv("user_password")
	dbName := os.Getenv("dbname")
	port := os.Getenv("postgres_port")

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable",
		user_name, user_password, dbName, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error initializing the postgres", err.Error())
		return nil, err
	}

	db.AutoMigrate(&models.Account{})

	return &PostgresAccountRepository{Db: db}, nil
}

func (r *PostgresAccountRepository) CreateAccount(ctx context.Context, acc *models.Account) error {
	if acc == nil {
		return errors.New("account cannot be nil")
	}

	tx := r.Db.WithContext(ctx).Save(&acc)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected < 1 {
		return errors.New("no rows affected while creating account")
	}

	return nil
}

func (r *PostgresAccountRepository) GetAccount(ctx context.Context, id string) (*models.Account, error) {
	var account models.Account
	tx := r.Db.WithContext(ctx).First(&account, "account_id = ?", id)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user account not exits")
		}
		return nil, tx.Error
	}

	return &account, nil
}

// UpdateAccountBalance updates the balance of an account atomically within a transaction.
func (r *PostgresAccountRepository) UpdateAccountBalance(ctx context.Context, accountID uint, newBalance float32) error {
	result := r.Db.WithContext(ctx).Model(&models.Account{}).Where("id = ?", accountID).Update("balance", newBalance)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return errors.New("no rows affected while updating balance")
	}
	return nil
}

func (r *PostgresAccountRepository) GetMaxID(ctx context.Context) (maxID int) {

	r.Db.WithContext(ctx).Model(&models.Account{}).Select("MAX(id)").Row().Scan(&maxID)

	return
}
