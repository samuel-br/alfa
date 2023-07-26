package db

import (
	"alfa/common"
	"time"

	"gorm.io/gorm"
)

type AdvanceService struct {
	client *gorm.DB
}

func (service AdvanceService) GetAdvancePayToPreform() ([]Advance, error) {
	var advancepayments []Advance
	result := service.client.Joins("JOIN transactions ON transactions.transaction_id = advances.transaction_id").
		Where("transactions.status = ? AND advances.next_pay_date = ? AND advances.transaction < 12", "success", time.Now().Format("2006-01-02")).
		Find(&advancepayments)

	if result.Error != nil {
		return advancepayments, nil
	}

	return advancepayments, nil
}

func (service AdvanceService) GetAdvanceById(transactionId string) (Advance, error) {

	advance := Advance{
		Transaction_id: transactionId,
	}

	result := service.client.First(&advance, "transaction_id = ?", advance.Transaction_id)
	if result.Error != nil {
		return Advance{}, result.Error
	}

	return advance, nil
}

func (service AdvanceService) SaveAdvance(transaction common.SaveAdvanceTransactionReq) (Advance, error) {

	advancedTransaction := Advance{
		Transaction_id:    transaction.Transaction_id,
		Dst_bank_account:  transaction.Dst_bank_account,
		Amount:            transaction.Amount,
		Time:              time.Now(),
		Debit_transaction: []string{},
		Transactions:      0,
		Next_Pay_date:     time.Now().AddDate(0, 0, 7).Format("2006-01-02"),
	}

	result := service.client.Create(&advancedTransaction)
	if result.Error != nil {
		return Advance{}, result.Error
	}

	return advancedTransaction, nil
}

func (service AdvanceService) UpdateAdvance(advanceTransaction common.UpdateAdvanceTransactionReq) error {

	advanced := Advance{
		Transaction_id: advanceTransaction.Transaction_id,
	}

	service.client.First(&advanced, "transaction_id = ?", advanced.Transaction_id)

	if advanceTransaction.Debit_transaction != "" {
		advanced.Debit_transaction = append(advanced.Debit_transaction, advanceTransaction.Debit_transaction)
	}

	result := service.client.Model(&Advance{Transaction_id: advanceTransaction.Transaction_id}).
		Updates(&Advance{Transactions: advanceTransaction.Transactions,
			Debit_transaction: advanced.Debit_transaction, Next_Pay_date: advanceTransaction.Next_Pay_date})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func NewAdvanceService(client *gorm.DB) AdvanceRepo {
	return AdvanceService{client: client}
}
