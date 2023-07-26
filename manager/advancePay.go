package manager

import (
	"alfa/common"
	"bytes"
	"encoding/json"
	"io"

	"net/http"
)

func (service *BillingService) AdvancePay(advanceTransactionReq common.AdvanceTransactionReq) error {

	transactionReq := common.TransactionReq{
		Src_bank_account: common.INTERNAL_BANK_ACCOUNT,
		Dst_bank_account: advanceTransactionReq.Dst_bank_account,
		Amount:           advanceTransactionReq.Amount,
		Direction:        "credit",
	}

	transactionid,err := service.PerformTranscation(transactionReq)
	
	if err != nil{
		return err
	}

	saveTransactionReq := common.SaveTransactionReq{
		Transaction_id:     transactionid,
		Src_bank_account:   common.INTERNAL_BANK_ACCOUNT,
		Dst_bank_account:   transactionReq.Dst_bank_account,
		Amount:             transactionReq.Amount,
		Direction:          "credit",
		Source_transaction: transactionid,
	}

	saveAdvanceTransactionReq := common.SaveAdvanceTransactionReq{
		Transaction_id:   transactionid,
		Dst_bank_account: transactionReq.Dst_bank_account,
		Amount:           transactionReq.Amount,
	}

	_, err = service.TransactionClient.SaveTransaction(saveTransactionReq)
	if err != nil {
		return err
	}

	_, err = service.Advanceclient.SaveAdvance(saveAdvanceTransactionReq)
	if err != nil {
		return err
	}
	return nil
}

func (service *BillingService) PerformTranscation(transactionReq common.TransactionReq) (string, error) {
	jsonData, err := json.Marshal(transactionReq)

	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", common.PERFORM_TRANSACTIO_URL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", err
	}
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	transactionid := string(responseBody)
	return transactionid, nil

}
