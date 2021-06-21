package domain

import (
	"github.com/dbielecki97/banking/dto"
	"github.com/dbielecki97/banking/errs"
)

type Account struct {
	AccountId   string `db:"account_id"`
	CustomerId  string `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	AccountType string `db:"account_type"`
	Amount      float64
	Status      string
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	FindById(string) (*Account, *errs.AppError)
	SaveTransaction(Transaction) (*Transaction, *errs.AppError)
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: a.AccountId}
}

func (a *Account) Update(transaction Transaction) {
	if transaction.TransactionType == "deposit" {
		a.Amount += transaction.Amount
	} else if transaction.TransactionType == "withdraw" {
		a.Amount -= transaction.Amount
	}
}

func (a Account) CanWithdraw(amount float64) bool {
	if a.Amount > amount {
		return true
	}

	return false
}
