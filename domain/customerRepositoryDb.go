package domain

import (
	"database/sql"
	"fmt"
	"github.com/dbielecki97/banking/errs"
	"github.com/dbielecki97/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"time"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func NewCustomerRepositoryDb() *CustomerRepositoryDb {
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
	return &CustomerRepositoryDb{client}
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllStatusSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllStatusSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpected("unexpected database error")
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var c Customer
	err := d.client.Get(&c, customerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFound("customer not found")
		}
		logger.Error("Error while scanning customer " + err.Error())
		return nil, errs.NewUnexpected("unexpected database error")
	}

	return &c, nil
}
