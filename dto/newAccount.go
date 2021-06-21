package dto

import (
	"github.com/dbielecki97/banking-lib/errs"
	"strings"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id,omitempty"`
	AccountType string  `json:"account_type,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidation("To open a new account you need to deposit at least 5000.00")
	}

	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errs.NewValidation("Account type should be checking or saving")
	}

	return nil
}

type NewAccountResponse struct {
	AccountId string `json:"account_id,omitempty"`
}
