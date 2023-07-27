package api

import (
	"alfa/common"
	"alfa/manager"
	"github.com/gin-gonic/gin"
)

type AdvencePayApiService struct {
	BillingService manager.BillingService
}

func (service AdvencePayApiService) AdvancePay(c *gin.Context) {

	var advanceTransactionReq common.AdvanceTransactionReq
	if err := c.BindJSON(&advanceTransactionReq); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := service.BillingService.AdvancePay(advanceTransactionReq)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.Status(200)
}

func NewAdvanceApiService() AdvencePayApiService {
	
	billingService := manager.NewDefaultBillingService()
	return AdvencePayApiService{
		BillingService: billingService,
	}
}
