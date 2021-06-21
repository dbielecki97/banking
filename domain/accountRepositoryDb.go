package domain

import (
	"database/sql"
	"github.com/dbielecki97/banking-lib/errs"
	"github.com/dbielecki97/banking-lib/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func NewAccountRepositoryDb(client *sqlx.DB) *AccountRepositoryDb {
	return &AccountRepositoryDb{client}
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values(?, ?, ?, ?, ?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpected("unexpected database error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id: " + err.Error())
		return nil, errs.NewUnexpected("unexpected database error")
	}

	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func (d AccountRepositoryDb) FindById(id string) (*Account, *errs.AppError) {
	customerSql := "select account_id, customer_id, opening_date, account_type, amount, status from accounts where account_id = ?"

	var c Account
	err := d.client.Get(&c, customerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFound("account not found")
		}
		logger.Error("Error while scanning account " + err.Error())
		return nil, errs.NewUnexpected("unexpected database error")
	}

	return &c, nil
}

func (d AccountRepositoryDb) Update(a *Account) *errs.AppError {
	updateSql := "UPDATE accounts set amount = ? WHERE account_id = ?"

	_, err := d.client.Exec(updateSql, a.Amount, a.AccountId)
	if err != nil {
		return errs.NewUnexpected("unexpected database error")
	}

	return nil
}

func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for a bank account transaction: " + err.Error())
		return nil, errs.NewUnexpected("unexpected database error")
	}

	insertSql := "INSERT INTO transactions(account_id, amount, transaction_type, transaction_date) values(?, ?, ?, ?)"
	result, _ := tx.Exec(insertSql, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	if t.IsWithdrawal() {
		_, err = tx.Exec("UPDATE accounts set amount = amount - ? WHERE account_id = ?", t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec("UPDATE accounts set amount = amount + ? WHERE account_id = ?", t.Amount, t.AccountId)
	}
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpected("unexpected database error")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while committing transaction for bank account: " + err.Error())
		return nil, errs.NewUnexpected("unexpected database error")
	}
	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last transaction id: " + err.Error())
		return nil, errs.NewUnexpected("unexpected database error")
	}

	account, appErr := d.FindById(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}

	t.TransactionId = strconv.FormatInt(transactionId, 10)
	t.Amount = account.Amount

	return &t, nil
}
