package dto

type TransactionResponse struct {
	TransactionId   string  `json:"transaction_id,omitempty"`
	NewBalance      float64 `json:"new_balance,omitempty"`
	AccountId       string  `json:"account_id,omitempty"`
	TransactionType string  `json:"transaction_type,omitempty"`
	TransactionDate string  `json:"transaction_date,omitempty"`
}
