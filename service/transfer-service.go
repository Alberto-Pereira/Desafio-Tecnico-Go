// Service package contains service operations
// for account and transfer models
package service

import (
	"desafio-tecnico/model"
	"desafio-tecnico/repository"
	"desafio-tecnico/security"
	"errors"
	"fmt"
	"time"
)

// Create Transfer
// Receives a token, account destination id and an amount
// then validates, update the accounts, create a transfer and send to the repository
// If the operation is successful, returns nil
// If the operation fails, returns an error
func CreateTransfer(token string, accountDestinationId int, amount int) error {

	accountOriginId, err := security.ValidateToken(token)
	if err != nil {
		return err
	}

	err = validateTransfer(accountOriginId, accountDestinationId, amount)
	if err != nil {
		return err
	}

	accountOriginBalance, err := validateAccountBalance(accountOriginId, amount)
	if err != nil {
		return err
	}

	err = updateAccounts(accountOriginId, accountOriginBalance, accountDestinationId, amount)
	if err != nil {
		return err
	}

	err = makeTransfer(accountOriginId, accountDestinationId, amount)
	if err != nil {
		return err
	}

	return nil
}

// Read Transfers
// Receives a token, validates and search for all the transfers of the account who requested
// If the operation is successful, returns the transfers and nil
// If the operation fails, returns nil and an error
func ReadTransfers(token string) ([]model.Transfer, error) {

	accountId, err := security.ValidateToken(token)
	if err != nil {
		return nil, err
	}

	err = repository.ReadAccountId(accountId)
	if err != nil {
		return nil, err
	}

	transfers, err := repository.ReadTransfers(accountId)
	if err != nil {
		return nil, err
	}

	return transfers, nil
}

// Validate Transfer
// Receives an account origin id, account destination id and an amount then validates
// If the operation is successful, returns nil
// If the operation fails, returns an error
func validateTransfer(accountOriginId int, accountDestinationId int, amount int) error {

	if accountOriginId == accountDestinationId {
		return errors.New("The destination account is the same as origin account!")
	}

	if amount <= 0 {
		return errors.New("Invalid amount to transfer!")
	}

	err := repository.ReadAccountId(accountDestinationId)
	if err != nil {
		return err
	}

	return nil
}

// Validate Account Balance
// Receives an account origin id and an amount, search for that account balance then validates
// If the operation is successful, returns the account origin balance and nil
// If the operation fails, returns 0 and an error
func validateAccountBalance(accountOriginId int, amount int) (int, error) {

	accountOriginBalance, err := repository.ReadAccountBalance(accountOriginId)
	if err != nil {
		return 0, err
	}

	accBalance := float64(accountOriginBalance)
	amt := float64(amount)

	if accountOriginBalance <= 0 {
		return 0, fmt.Errorf("Invalid account balance! You have %.2f", accBalance/100)
	}

	if accountOriginBalance < amount {
		return 0, fmt.Errorf("Invalid amount to transfer! You have %.2f and wants to transfer %.2f", accBalance/100, amt/100)
	}

	return accountOriginBalance, nil
}

// Update Accounts
// Receives an account origin id and balance, account destination id and amount to transfer
// then update the accounts based in the difference
// If the operation is successful, returns nil
// If the operation fails, returns an error
func updateAccounts(accountOriginId int, accountOriginBalance int, accountDestinationId int, amount int) error {

	newAccountOriginBalance := accountOriginBalance - amount

	err := repository.UpdateAccountBalance(accountOriginId, newAccountOriginBalance)
	if err != nil {
		return err
	}

	accountDestinationBalance, err := repository.ReadAccountBalance(accountDestinationId)
	if err != nil {
		return err
	}

	newAccountDestinationBalance := accountDestinationBalance + amount

	err = repository.UpdateAccountBalance(accountDestinationId, newAccountDestinationBalance)
	if err != nil {
		return err
	}

	return nil
}

// Make Transfer
// Receives an account origin id, account destination id and an amount
// then create a transfer and send to the repository
// If the operation is successful, returns nil
// If the operation fails, returns an error
func makeTransfer(accountOriginId int, accountDestinationId int, amount int) error {

	createdAt := time.Now().Unix()

	transfer := model.Transfer{Account_origin_id: accountOriginId, Account_destination_id: accountDestinationId,
		Amount: amount, Created_at: int(createdAt)}

	err := repository.CreateTransfer(transfer)
	if err != nil {
		return err
	}

	return nil
}
