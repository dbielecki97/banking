package app

import (
	"github.com/dbielecki97/banking/domain"
	"github.com/dbielecki97/banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	ch := CustomerHandler{service.NewDefaultCustomerService(domain.NewCustomerRepositoryStub())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}
