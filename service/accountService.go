package service

import (
	"github.com/dbielecki97/banking/domain"
	"github.com/dbielecki97/banking/dto"
	"github.com/dbielecki97/banking/errs"
	"time"
)

const TSLayout = "2006-01-02 15:04:05"

//go:generate mockgen -destination=../mocks/service/mockAccountService.go -package=service github.com/dbielecki97/banking/service AccountService
type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	accountRepo domain.AccountRepository
}

func NewDefaultAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{accountRepo: repo}
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	newAccount, err := s.accountRepo.Save(a)
	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAccountResponseDto()
	return &response, nil
}

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	if req.IsTransactionTypeWithdrawal() {
		acc, err := s.accountRepo.FindById(req.AccountId)
		if err != nil {
			return nil, err
		}

		if !acc.CanWithdraw(req.Amount) {
			return nil, errs.NewValidation("Insufficient balance in the account")
		}
	}

	t := domain.Transaction{
		TransactionId:   "",
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format(TSLayout),
	}

	transaction, err := s.accountRepo.SaveTransaction(t)
	if err != nil {
		return nil, err
	}

	response := transaction.ToDto()

	return &response, nil
}
