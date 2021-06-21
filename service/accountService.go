package service

import (
	"github.com/dbielecki97/banking/domain"
	"github.com/dbielecki97/banking/dto"
	"github.com/dbielecki97/banking/errs"
)

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
	if err := req.Validate(); err != nil {
		return nil, err
	}

	account := domain.NewAccount(req)
	if newAccount, err := s.accountRepo.Save(account); err != nil {
		return nil, err
	} else {
		return newAccount.ToNewAccountResponseDto(), nil
	}
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

	transaction := domain.NewTransaction(req)
	newTransaction, err := s.accountRepo.SaveTransaction(transaction)
	if err != nil {
		return nil, err
	}

	response := newTransaction.ToDto()

	return &response, nil
}
