package app

import (
	"fmt"
	"github.com/dbielecki97/banking/domain"
	"github.com/dbielecki97/banking/logger"
	"github.com/dbielecki97/banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func Start() {
	router := mux.NewRouter()

	sanityCheck()

	//ch := CustomerHandler{service.NewDefaultCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandler{service.NewDefaultCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router)
	if err != nil {
		log.Fatal(err)
	}
}

func sanityCheck() {
	keys := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME"}

	allPresent := true
	for _, e := range keys {
		ok := checkEnvVariable(e)
		if allPresent != false {
			allPresent = ok
		}
	}

	if !allPresent {
		os.Exit(1)
	}
}

func checkEnvVariable(key string) bool {
	if os.Getenv(key) == "" {
		logger.Error("Environment variable " + key + " not defined!")
		return false
	}
	return true
}
