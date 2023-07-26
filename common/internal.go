package common

const INTERNAL_BANK_ACCOUNT = "12345"

const PERFORM_TRANSACTIO_URL =  "https://blackbox/api/perform-transaction"
const DONLOAD_REPORT_URL = "https://blackbox/download_report"

type TransactionReq struct {
	Src_bank_account string `json:"src_bank_account"`
	Dst_bank_account string `json:"dst_bank_account"`
	Amount           float64    `json:"amount"`
	Direction        string `json:"direction"`
}

type SaveTransactionReq struct {
	Src_bank_account   string `json:"src_bank_account"`
	Dst_bank_account   string `json:"dst_bank_account"`
	Amount             float64    `json:"amount"`
	Direction          string `json:"direction"`
	Transaction_id     string `json:"transaction_id"`
	Source_transaction string `json:"Source_transaction"`
}

type SaveAdvanceTransactionReq struct {
	Transaction_id   string `json:"transaction_id"`
	Dst_bank_account string `json:"dst_bank_account"`
	Amount           float64    `json:"amount"`
}

type UpdateTransactionReq struct {
	Transaction_id string `json:"transaction_id"`
	Status         string `json:"status"`
}

type AdvanceTransactionReq struct {
	Dst_bank_account string `json:"dst_bank_account"`
	Amount           float64    `json:"amount"`
}

type UpdateAdvanceTransactionReq struct {
	Transaction_id    string `json:"transaction_id"`
	Debit_transaction string `json:"debit_transaction"`
	Transactions      int    `json:"transactions"`
	Next_Pay_date     string `json:"next_pay_date"`
}
