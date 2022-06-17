package repository

import (
	"desafio-tecnico/model"
	"errors"
)

func CreateAccount(account model.Account) error {

	db := StartDB()

	sqlStatement := `INSERT INTO desafiotecnicoprincipal.accounts(name, cpf, secret, balance, created_at) 
						VALUES ($1, $2, $3, $4, $5) RETURNING id;`

	var accId int

	err := db.QueryRow(sqlStatement, account.Name, account.CPF, account.Secret, account.Balance, account.Created_at).Scan(&accId)
	if err != nil || accId == 0 {
		return errors.New("Error while try to create account!")
	}

	return nil
}

func ReadAccount(accountCpf string) (int, string, error) {

	db := StartDB()

	sqlStatement := `SELECT id, secret FROM desafiotecnicoprincipal.accounts WHERE cpf=$1;`

	var accId int
	var accSecret string

	err := db.QueryRow(sqlStatement, accountCpf).Scan(&accId, &accSecret)
	if err != nil || accId == 0 {
		return 0, "", errors.New("Invalid account cpf!")
	}

	return accId, accSecret, nil
}

func ReadAccountId(accountId int) error {

	db := StartDB()

	sqlStatement := `SELECT id FROM desafiotecnicoprincipal.accounts WHERE id=$1`

	var accId int

	err := db.QueryRow(sqlStatement, accountId).Scan(&accId)
	if err != nil || accountId != accId {
		return errors.New("Account id not found!")
	}

	return nil
}

func ReadAccountCpf(accountCpf string) (bool, error) {

	db := StartDB()

	sqlStatement := `SELECT cpf FROM desafiotecnicoprincipal.accounts WHERE cpf=$1;`

	var accCpf string

	err := db.QueryRow(sqlStatement, accountCpf).Scan(&accCpf)
	if err != nil && accCpf != "" {
		return false, errors.New("Error while try to search for account cpf!")
	}

	if accCpf == "" {
		return false, nil
	}

	return true, nil
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

func UpdateAccountBalance(accountId int, accountBalance int) error {

	db := StartDB()

	sqlStatement := `UPDATE desafiotecnicoprincipal.accounts SET balance=$2 WHERE id=$1 RETURNING id;`

	var accId int

	err := db.QueryRow(sqlStatement, accountId, accountBalance).Scan(&accId)
	if err != nil || accId != accountId {
		return errors.New("Error while try to update account balance!")
	}

	return nil
}
