package dto

import (
	"net/http"
	"testing"
)

func Test_should_return_error_when_transaction_type_is_not_deposit_or_withdrawal(t *testing.T) {
	request := TransactionRequest{TransactionType: "invalid"}

	err := request.Validate()

	if err.Message != "Available transaction types are only deposit and withdraw" {
		t.Error("Invalid message while validating transaction type")
	}

	if err.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while validating transaction type")
	}
}

func Test_should_return_error_when_amount_is_less_than_zero(t *testing.T) {
	request := TransactionRequest{Amount: -500, TransactionType: WITHDRAWAL}

	err := request.Validate()

	if err.Message != "Amount cannot be negative" {
		t.Error("Invalid message while validating transaction amount")
	}

	if err.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while validating transaction amount")
	}
}
