package app

import (
	"github.com/dbielecki97/banking-lib/errs"
	"github.com/dbielecki97/banking/dto"
	"github.com/dbielecki97/banking/mocks/service"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *mux.Router
var ch CustomerHandler
var ah AccountHandler
var mockCustomerService *service.MockCustomerService
var mockAccountService *service.MockAccountService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockCustomerService = service.NewMockCustomerService(ctrl)
	mockAccountService = service.NewMockAccountService(ctrl)

	ch = CustomerHandler{service: mockCustomerService}
	ah = AccountHandler{service: mockAccountService}

	router = mux.NewRouter()

	router.HandleFunc("/customers", ch.getAllCustomers)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer)

	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.newAccount)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.newTransaction)

	return func() {
		defer ctrl.Finish()
	}
}

func Test_should_return_customers_with_status_code_200(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	dummyCustomers := []dto.CustomerResponse{
		{
			Id:          "1001",
			Name:        "Dawid",
			City:        "Warsaw",
			Zipcode:     "15-222",
			DateOfBirth: "2000-01-01",
			Status:      "1",
		}, {
			Id:          "1002",
			Name:        "Michał",
			City:        "Warsaw",
			Zipcode:     "111111",
			DateOfBirth: "2000-01-01",
			Status:      "1",
		},
	}
	mockCustomerService.EXPECT().GetAllCustomers("").Return(dummyCustomers, nil)

	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_status_code_500_with_error_message(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	mockCustomerService.EXPECT().GetAllCustomers("").Return(nil, errs.NewUnexpected("some database error"))

	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_customer_with_status_code_200(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	dummyCustomer := dto.CustomerResponse{

		Id:          "1001",
		Name:        "Dawid",
		City:        "Warsaw",
		Zipcode:     "15-222",
		DateOfBirth: "2000-01-01",
		Status:      "1",
	}

	mockCustomerService.EXPECT().GetCustomer("2000").Return(&dummyCustomer, nil)

	request, _ := http.NewRequest(http.MethodGet, "/customers/2000", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_status_code_500_with_error_message_when_getting_customer(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockCustomerService.EXPECT().GetCustomer("2000").Return(nil, errs.NewUnexpected("some database error"))

	request, _ := http.NewRequest(http.MethodGet, "/customers/2000", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_status_code_404_with_error_message(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockCustomerService.EXPECT().GetCustomer("2000").Return(nil, errs.NewNotFound("customer not found"))

	request, _ := http.NewRequest(http.MethodGet, "/customers/2000", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusNotFound {
		t.Error("Failed while testing the status code")
	}
}
