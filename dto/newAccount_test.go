package dto

import (
	"net/http"
	"testing"
)

func Test_should_return_error_when_creating_account_with_too_little_deposit(t *testing.T) {
	request := NewAccountRequest{
		AccountType: "saving",
		Amount:      500,
	}

	err := request.Validate()

	if err.Message != "To open a new account you need to deposit at least 5000.00" {
		t.Error("Invalid message while validating amount")
	}

	if err.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while validating amount")
	}
}

func Test_should_return_error_when_creating_account_using_invalid_type(t *testing.T) {
	request := NewAccountRequest{
		AccountType: "savingg",
		Amount:      5000,
	}

	err := request.Validate()

	if err.Message != "Account type should be checking or saving" {
		t.Error("Invalid message while validating account type")
	}

	if err.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while validating account type")
	}
}
