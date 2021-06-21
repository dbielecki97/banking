package domain

import (
	"github.com/dbielecki97/banking/dto"
	"time"
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

func NewTransaction(req dto.TransactionRequest) Transaction {
	return Transaction{
		TransactionId:   "",
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format(TSLayout),
	}
}
