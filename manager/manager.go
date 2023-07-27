package manager

import (
	"alfa/common"
	"alfa/db"
	"alfa/utils"
)

type DefaultBillingService struct {
	Advanceclient     db.AdvanceRepo
	TransactionClient db.TransactionRepo
}

type BillingService interface {
	AdvancePay(advanceTransactionReq common.AdvanceTransactionReq) error
	PerformTranscation(transactionReq common.TransactionReq) (string, error)
	Update() error
	GetData() (TransactionDataFileForm, error)
	nextPayDateCal() string
	DebitPay() error
}

func NewDefaultBillingService() BillingService {
	advanceClient, _ := utils.ConnectDB()
	transactionClient, _ := utils.ConnectDB()

	defaultBillingService := DefaultBillingService{
		Advanceclient:     db.NewAdvanceService(advanceClient),
		TransactionClient: db.NewTransactionService(transactionClient),
	}
	return &defaultBillingService
}
