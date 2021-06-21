package domain

import (
	"github.com/dbielecki97/banking/dto"
)

const WITHDRAWAL = "withdrawal"

type Transaction struct {
	TransactionId   string
	AccountId       string
	Amount          float64
	TransactionType string
	TransactionDate string
}

type TransactionRepository interface {
}

func (t Transaction) IsWithdrawal() bool {
	if t.TransactionType == WITHDRAWAL {
		return true
	}
	return false
}

func (t Transaction) ToDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		TransactionId:   t.TransactionId,
		NewBalance:      t.Amount,
		AccountId:       t.AccountId,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}
