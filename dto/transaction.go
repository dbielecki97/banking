package dto

import (
	"github.com/dbielecki97/banking-lib/errs"
)

type TransactionRequest struct {
	AccountId       string  `json:"account_id,omitempty"`
	TransactionType string  `json:"transaction_type,omitempty"`
	Amount          float64 `json:"amount,omitempty"`
}

const (
	WITHDRAWAL = "withdrawal"
	DEPOSIT    = "deposit"
)

func (r TransactionRequest) Validate() *errs.AppError {
	if !r.IsTransactionTypeWithdrawal() && !r.IsTransactionTypeDeposit() {
		return errs.NewValidation("Available transaction types are only deposit and withdraw")
	}

	if r.Amount < 0 {
		return errs.NewValidation("Amount cannot be negative")
	}

	return nil
}

func (r TransactionRequest) IsTransactionTypeWithdrawal() bool {
	return r.TransactionType == WITHDRAWAL
}

func (r TransactionRequest) IsTransactionTypeDeposit() bool {
	return r.TransactionType == DEPOSIT
}

type TransactionResponse struct {
	TransactionId   string  `json:"transaction_id,omitempty"`
	NewBalance      float64 `json:"new_balance,omitempty"`
	AccountId       string  `json:"account_id,omitempty"`
	TransactionType string  `json:"transaction_type,omitempty"`
	TransactionDate string  `json:"transaction_date,omitempty"`
}
