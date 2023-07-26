package main

import (
	"alfa/db"
	"alfa/utils"
	"alfa/manager"
	"time"
)

func main() {
	advanceClient, _ := utils.ConnectDB()
	transactionClient, _ := utils.ConnectDB()

	billingService := manager.BillingService{
		Advanceclient:    db.NewAdvanceService(advanceClient),
		TransactionClient: db.NewTransactionService(transactionClient),
	
	}
	interval := 12 * time.Hour
	ticker := time.NewTicker(interval)

	for range ticker.C {
		billingService.Update()
		billingService.DebitPay()
	}
}

