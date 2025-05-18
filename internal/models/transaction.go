package models

import "time"

type Transaction struct {
	AccountID   string          `json:"account_id" bson:"account_id"`
	Amount      float64         `json:"amount" bson:"amount"`
	Type        TransactionType `json:"type" bson:"type"` //Credit or Debit
	Timestamp   time.Time       `json:"time_stamp" bson:"timestamp" `
	Description string          `json:"description" bson:"description"`
}

type TransactionType string

const (
	Credit TransactionType = "credit"
	Debit  TransactionType = "debit"
)

type Date struct {
	FromDate string `json:"start_date"`
	ToDate   string `json:"end_date"`
}
