package app

import (
	"encoding/json"
	"github.com/dbielecki97/banking/dto"
	"github.com/dbielecki97/banking/errs"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_should_create_new_account(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	response := dto.NewAccountResponse{
		AccountId: "123123",
	}
	expect := dto.NewAccountRequest{
		CustomerId:  "2000",
		AccountType: "saving",
		Amount:      5000,
	}

	mockAccountService.EXPECT().NewAccount(expect).Return(&response, nil)

	r := dto.NewAccountRequest{
		CustomerId:  "",
		AccountType: "saving",
		Amount:      5000,
	}

	marshal, _ := json.Marshal(r)
	reader := strings.NewReader(string(marshal))

	request, _ := http.NewRequest(http.MethodPost, "/customers/2000/account", reader)
	request.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusCreated {
		t.Error("Failed while testing the status code")
	}

	err := json.NewDecoder(recorder.Body).Decode(&response)
	if err != nil {
		t.Error(err.Error())
	}

	if response.AccountId != "123123" {
		t.Error("invalid account id returned")
	}
}

func Test_should_not_pass_type_validation_and_return_error_with_message(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	expect := dto.NewAccountRequest{
		CustomerId:  "2000",
		AccountType: "savingg",
		Amount:      5000,
	}

	mockAccountService.EXPECT().NewAccount(expect).Return(nil, errs.NewValidation("Account type should be checking or saving"))

	r := dto.NewAccountRequest{
		CustomerId:  "",
		AccountType: "savingg",
		Amount:      5000,
	}

	marshal, _ := json.Marshal(r)
	reader := strings.NewReader(string(marshal))

	request, _ := http.NewRequest(http.MethodPost, "/customers/2000/account", reader)
	request.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusUnprocessableEntity {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_not_pass_amount_validation_and_return_error_with_message(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	expect := dto.NewAccountRequest{
		CustomerId:  "2000",
		AccountType: "saving",
		Amount:      4000,
	}

	mockAccountService.EXPECT().NewAccount(expect).Return(nil, errs.NewValidation("To open a new account you need to deposit at least 5000.00"))

	r := dto.NewAccountRequest{
		CustomerId:  "",
		AccountType: "saving",
		Amount:      4000,
	}

	marshal, _ := json.Marshal(r)
	reader := strings.NewReader(string(marshal))

	request, _ := http.NewRequest(http.MethodPost, "/customers/2000/account", reader)
	request.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusUnprocessableEntity {
		t.Error("Failed while testing the status code")
	}
}
