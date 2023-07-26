package manager

import (
	"alfa/common"
	"bufio"
	"io"
	"net/http"
	"os"
	"strings"
)

var filepath = os.Getenv("DATA_FILE_PATH")

type TransactionDataFileForm = map[string]string

func (service *BillingService) Update() error {

	data, err := service.GetData()
	if err != nil {
		return err
	}

	for transactionId, status := range data {
		transaction, _ := service.TransactionClient.GetTransactionById(transactionId)
		if transaction.Status == "pending" {
			updaterequest := common.UpdateTransactionReq{
				Transaction_id: transactionId,
				Status:         status,
			}
			service.TransactionClient.UpdateTransaction(updaterequest)

			if transaction.Transaction_id != transaction.Source_transaction {
				updateAdvance := common.UpdateAdvanceTransactionReq{
					Transaction_id:    transaction.Source_transaction,
					Debit_transaction: transaction.Transaction_id,
				}
				advanceTransaction, _ := service.Advanceclient.GetAdvanceById(transaction.Source_transaction)

				if transaction.Status == "success" {
					updateAdvance.Transactions = advanceTransaction.Transactions + 1
				}
				service.Advanceclient.UpdateAdvance(updateAdvance)
			}
		}

	}
	return nil
}

func (b BillingService) GetData() (TransactionDataFileForm, error) {
	err := DownloadFile()
	if err != nil {
		return TransactionDataFileForm{}, err
	}

	file, err := os.Open(filepath)
	if err != nil {
		return TransactionDataFileForm{}, err
	}
	defer file.Close()

	transactionDataFileForm := TransactionDataFileForm{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")

		transactionID := fields[0]
		status := fields[1]
		transactionDataFileForm[transactionID] = status
	}

	if err = scanner.Err(); err != nil {
		return TransactionDataFileForm{}, err
	}

	return transactionDataFileForm, nil

}

func DownloadFile() error {

	resp, err := http.Get(common.DONLOAD_REPORT_URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
