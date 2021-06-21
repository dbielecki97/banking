package dto

import "github.com/dbielecki97/banking/errs"

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
	if r.TransactionType != WITHDRAWAL && r.TransactionType != DEPOSIT {
		return errs.NewValidation("Available transaction types are only deposit and withdraw")
	}

	if r.Amount < 0 {
		return errs.NewValidation("Amount cannot be negative")
	}

	return nil
}

func (r TransactionRequest) InTransactionTypeWithdrawal() bool {
	if r.TransactionType == WITHDRAWAL {
		return true
	}

	return false
}
