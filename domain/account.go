package domain

import (
	"github.com/dbielecki97/banking/dto"
	"github.com/dbielecki97/banking/errs"
	"time"
)

const TSLayout = "2006-01-02 15:04:05"

type Account struct {
	AccountId   string `db:"account_id"`
	CustomerId  string `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	AccountType string `db:"account_type"`
	Amount      float64
	Status      string
}

//go:generate mockgen -destination=../mocks/domain/mockAccountRepository.go -package=domain github.com/dbielecki97/banking/domain AccountRepository
type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	FindById(string) (*Account, *errs.AppError)
	SaveTransaction(Transaction) (*Transaction, *errs.AppError)
}

func (a Account) ToNewAccountResponseDto() *dto.NewAccountResponse {
	return &dto.NewAccountResponse{AccountId: a.AccountId}
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

func NewAccount(req dto.NewAccountRequest) Account {
	return Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format(TSLayout),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
}
