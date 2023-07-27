package manager

import (
	"alfa/common"
	"time"
)

func (service *DefaultBillingService) DebitPay() error {
	
	advancePayToPreform, _ := service.Advanceclient.GetAdvancePayToPreform()

	for _, payToPreform := range advancePayToPreform {
		
		transactionReq := common.TransactionReq{
			Src_bank_account: common.INTERNAL_BANK_ACCOUNT,
			Dst_bank_account: payToPreform.Dst_bank_account,
			Amount:           payToPreform.Amount / 12,
			Direction:        "debit",
		}
	
		transactionId, err := service.PerformTranscation(transactionReq)
		
		if err != nil {
			return err
		}

		transactionToSave := common.SaveTransactionReq{
			Transaction_id:     transactionId,
			Amount:             transactionReq.Amount,
			Dst_bank_account:   transactionReq.Dst_bank_account,
			Src_bank_account:   transactionReq.Src_bank_account,
			Source_transaction: payToPreform.Transaction_id,
			Direction:          "debit",
		}
		service.TransactionClient.SaveTransaction(transactionToSave)

		service.Advanceclient.UpdateAdvance(common.UpdateAdvanceTransactionReq{
			Next_Pay_date:  service.nextPayDateCal(),
			Transaction_id: payToPreform.Transaction_id})
	}
	return nil
}

func (b DefaultBillingService) nextPayDateCal() string {
	return time.Now().AddDate(0, 0, 7).Format("2006-01-02")
}
