package repository

import (
	"desafio-tecnico/model"
	"errors"
)

func ReadAccounts() ([]model.Account, error) {

	db := StartDB()

	sqlStatement := `SELECT id, name, cpf, secret, balance, created_at
						FROM desafiotecnicoprincipal.accounts;`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		return nil, errors.New("Error while try to read accounts!")
	}

	var account model.Account
	var accounts []model.Account

	for rows.Next() {

		err := rows.Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.Created_at)

		if err != nil {
			return nil, errors.New("Error while try to scan a row of accounts!")
		}

		accounts = append(accounts, account)
	}

	err = rows.Err()

	if err != nil {
		return nil, errors.New("Unexpected error in accounts row!")
	}

	return accounts, nil
}

func ReadAccountBalance(accountId int) (int, error) {

	db := StartDB()

	sqlStatement := `SELECT id, balance FROM desafiotecnicoprincipal.accounts WHERE id=$1;`

	var accountBalance int
	var accId int

	err := db.QueryRow(sqlStatement, accountId).Scan(&accId, &accountBalance)

	if accountId != accId {
		return 0, errors.New("This account doesn't exist!")
	}

	if err != nil {
		return 0, errors.New("Error while try to retrieve balance from account!")
	}

	return accountBalance, nil
}

func CreateAccount(account model.Account) error {

	db := StartDB()

	sqlStatement := `INSERT INTO desafiotecnicoprincipal.accounts(name, cpf, secret, balance, created_at) 
						VALUES ($1, $2, $3, $4, $5) RETURNING id;`

	var id int

	err := db.QueryRow(sqlStatement, account.Name, account.CPF, account.Secret, account.Balance, account.Created_at).Scan(&id)

	if err != nil || id == 0 {
		return errors.New("Error while try to create account!")
	}

	return nil
}
