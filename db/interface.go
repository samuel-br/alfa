package db

import "alfa/common"

type TransactionRepo interface {
	GetTransactionById(transactionId string) (Transaction, error)
	SaveTransaction(transaction common.SaveTransactionReq) (Transaction, error)
	UpdateTransaction(transaction common.UpdateTransactionReq) error
}

type AdvanceRepo interface {
	GetAdvanceById(transactionId string) (Advance, error)
	SaveAdvance(transaction common.SaveAdvanceTransactionReq) (Advance, error)
	UpdateAdvance(transactionId common.UpdateAdvanceTransactionReq) error
	GetAdvancePayToPreform() ([]Advance, error)
}
