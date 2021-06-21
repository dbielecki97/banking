package service

import (
	"github.com/dbielecki97/banking-lib/errs"
	realDomain "github.com/dbielecki97/banking/domain"
	"github.com/dbielecki97/banking/dto"
	"github.com/dbielecki97/banking/mocks/domain"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

var mockRepo *domain.MockAccountRepository
var service AccountService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockRepo = domain.NewMockAccountRepository(ctrl)
	service = NewDefaultAccountService(mockRepo)

	return func() {
		service = nil
		defer ctrl.Finish()
	}
}

func Test_should_receive_error_from_validate_method(t *testing.T) {
	request := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      0,
	}

	service := NewDefaultAccountService(nil)

	_, appError := service.NewAccount(request)

	if appError == nil {
		t.Error("failed while testing the new account validation")
	}
}

func Test_should_receive_error_from_repository_when_db_cant_create_account(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	req := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      5000,
	}

	a := realDomain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	mockRepo.EXPECT().Save(a).Return(nil, errs.NewUnexpected("unexpected database error"))

	_, appError := service.NewAccount(req)

	if appError == nil {
		t.Error("Test failed while validating error for new account")
	}
}

func Test_should_return_new_account_response_when_a_new_account_is_saved_successfully(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	req := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      5000,
	}

	a := realDomain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	accountWithId := a
	accountWithId.AccountId = "201"
	mockRepo.EXPECT().Save(a).Return(&accountWithId, nil)

	newAccount, appError := service.NewAccount(req)

	if appError != nil {
		t.Error("Test failed while creating new error for new account")
	}

	if newAccount.AccountId != accountWithId.AccountId {
		t.Error("Failed while matching new account id")
	}
}
