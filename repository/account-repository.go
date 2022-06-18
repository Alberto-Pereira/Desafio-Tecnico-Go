// Repository package contains repository operations
// for account and transfer models
package repository

import (
	"desafio-tecnico/model"
	"errors"
)

// Create Account
// Receives an account and insert into the database
// If the operation is successful, returns nil
// If the operation fails, returns an error
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

// Read Account
// Receives an account cpf and searches for the account id and secret
// If the operation is successful, returns the account id, secret and nil
// If the operation fails, returns 0, "" and an error
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

// Read Account Id
// Receives an account id and search for the account id
// If the operation is successful, returns nil
// If the operation fails, returns an error
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

// Read Account Cpf
// Receives an account cpf and verify if the account cpf exists
// If the operation is successful, returns nil for an existing cpf
// If the operation fails, returns an error for a non  existing cpf or error
func ReadAccountCpf(accountCpf string) error {

	db := StartDB()

	sqlStatement := `SELECT cpf FROM desafiotecnicoprincipal.accounts WHERE cpf=$1;`

	var accCpf string

	err := db.QueryRow(sqlStatement, accountCpf).Scan(&accCpf)
	if err != nil {
		return errors.New("Error while try to search for account cpf!")
	}

	return nil
}

// Read Account Balance
// Receives an account id and search for the account balance
// If the operation is successful, returns the account balance and nil
// If the operation fails, returns 0 and an error
func ReadAccountBalance(accountId int) (int, error) {

	db := StartDB()

	sqlStatement := `SELECT id, balance FROM desafiotecnicoprincipal.accounts WHERE id=$1;`

	var accountBalance int
	var accId int

	err := db.QueryRow(sqlStatement, accountId).Scan(&accId, &accountBalance)
	if err != nil {
		return 0, errors.New("This account doesn't exist!")
	}

	return accountBalance, nil
}

// Read Accounts
// Searches for all the accounts into the database
// If the operation is successful, returns the accounts and nil
// If the operation fails, returns nil and and an error
func ReadAccounts() ([]model.Account, error) {

	db := StartDB()

	sqlStatement := `SELECT id, name, cpf, secret, balance, created_at
						FROM desafiotecnicoprincipal.accounts;`

	rows, err := db.Query(sqlStatement)
	if err != nil || rows.Next() == false {
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

// Update Account Balance
// Receives an account id and new balance value and updates the account
// If the operation is successful, returns nil
// If the operation fails, returns an error
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
