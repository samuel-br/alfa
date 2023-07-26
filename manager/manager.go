package manager

import "alfa/db"

type BillingService struct {
	Advanceclient     db.AdvanceRepo
	TransactionClient db.TransactionRepo
}

