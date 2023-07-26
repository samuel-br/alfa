package db

import (
	"alfa/common"
	"time"

	"gorm.io/gorm"
)

type TransactionService struct {
	client *gorm.DB
}

func (service TransactionService) GetTransactionById(transactionId string) (Transaction, error) {
	var transaction Transaction

	result := service.client.First(&transaction, "transaction_id = ?", transactionId)
	if result.Error != nil {
		return Transaction{}, result.Error
	}
	return transaction, nil
}

func (service TransactionService) SaveTransaction(transaction common.SaveTransactionReq) (Transaction, error) {
	transactionToSave := Transaction{
		Transaction_id:     transaction.Transaction_id,
		Src_bank_account:   transaction.Src_bank_account,
		Dst_bank_account:   transaction.Dst_bank_account,
		Amount:             transaction.Amount,
		Time:               time.Now(),
		Direction:          transaction.Direction,
		Source_transaction: transaction.Source_transaction,
	}

	result := service.client.Create(&transactionToSave)

	if result.Error != nil {
		return Transaction{}, result.Error
	}

	return transactionToSave, nil
}

func (service TransactionService) UpdateTransaction(updateTransactionReq common.UpdateTransactionReq) error {

	result := service.client.Model(&Transaction{Transaction_id: updateTransactionReq.Transaction_id}).
		Update("status", updateTransactionReq.Status)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func NewTransactionService(client *gorm.DB) TransactionRepo {
	return TransactionService{client: client}
}
