package app

import (
	"encoding/json"
	"github.com/dbielecki97/banking/dto"
	"github.com/dbielecki97/banking/service"
	"github.com/gorilla/mux"
	"net/http"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) newAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, err := h.service.NewAccount(request)
		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}

func (h AccountHandler) newTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["account_id"]

	var t dto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		t.AccountId = accountId
		response, err := h.service.MakeTransaction(t)
		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, response)
		}
	}
}
