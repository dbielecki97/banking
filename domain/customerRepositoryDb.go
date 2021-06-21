package domain

import (
	"database/sql"
	"github.com/dbielecki97/banking-lib/errs"
	"github.com/dbielecki97/banking-lib/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func NewCustomerRepositoryDb(client *sqlx.DB) *CustomerRepositoryDb {
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
