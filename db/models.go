package db

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Transaction_id     string `gorm:"primary_key;not null"`
	Src_bank_account   string `gorm:"not null"`
	Dst_bank_account   string `gorm:"not null"`
	Status             string
	Amount             float64       `gorm:"not null"`
	Time               time.Time `gorm:"not null"`
	Source_transaction string
	Direction          string
}

type Advance struct {
	gorm.Model
	Transaction_id    string    `gorm:"primary_key;not null"`
	Dst_bank_account  string    `gorm:"not null"`
	Amount            float64       `gorm:"not null"`
	Time              time.Time `gorm:"not null"`
	Debit_transaction []string
	Transactions      int
	Next_Pay_date     string
}
