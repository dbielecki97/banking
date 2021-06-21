package app

import (
	"fmt"
	"github.com/dbielecki97/banking/domain"
	"github.com/dbielecki97/banking/logger"
	"github.com/dbielecki97/banking/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"time"
)

func Start() {
	router := mux.NewRouter()

	sanityCheck()

	//ch := CustomerHandler{service.NewDefaultCustomerService(domain.NewCustomerRepositoryStub())}
	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	ch := CustomerHandler{service.NewDefaultCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewDefaultAccountService(accountRepositoryDb)}

	router.
		HandleFunc("/customers", ch.getAllCustomers).
		Methods(http.MethodGet).
		Name("GetAllCustomers")
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).
		Methods(http.MethodGet).
		Name("GetCustomer")
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.newAccount).
		Methods(http.MethodPost).
		Name("NewAccount")
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.newTransaction).
		Methods(http.MethodPost).
		Name("NewTransaction")

	am := AuthMiddleware{repo: domain.NewRemoteAuthRepository()}
	router.Use(am.authorizationHandler())

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

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := "%s:%s@tcp(%s:%s)/%s"
	client, err := sqlx.Open("mysql", fmt.Sprintf(dataSource, dbUser, dbPasswd, dbAddr, dbPort, dbName))
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxIdleConns(10)
	client.SetMaxOpenConns(10)
	return client
}
