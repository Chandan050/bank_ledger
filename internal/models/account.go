package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	AccountId   string  `json:"account_id"`
	Name        string  `json:"name"`
	Balance     float32 `json:"balance"`
	Address     string  `json:"address"`
	PhoneNumber int     `json:"phone_number"`
}
