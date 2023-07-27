package main

import (
	"alfa/manager"
	"time"
)

func main() {
	
	billingService := manager.NewDefaultBillingService()
	interval := 12 * time.Hour
	ticker := time.NewTicker(interval)

	for range ticker.C {
		billingService.Update()
		billingService.DebitPay()
	}
}

