package service

import (
	"desafio-tecnico/model"
	"desafio-tecnico/repository"
	"desafio-tecnico/security"
	"errors"
	"fmt"
	"time"
)

func ReadTransfers(token string) ([]model.Transfer, error) {

	accountId, err := security.ValidateToken(token)
	if err != nil {
		return nil, err
	}

	transfers, err := repository.ReadTransfers(accountId)
	if err != nil {
		return nil, err
	}

	if transfers == nil {
		return transfers, errors.New("This account doesn't have transfers!")
	}

	return transfers, nil
}

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

func validateAccountBalance(accountOriginId int, amount int) (int, error) {

	accountOriginBalance, err := repository.ReadAccountBalance(accountOriginId)
	if err != nil {
		return 0, err
	}

	accBalance := float64(accountOriginBalance)
	amt := float64(amount)

	if accountOriginBalance <= 0 {
		return 0, fmt.Errorf("Invalid account balance! You have %2.f", accBalance/100)
	}

	if accountOriginBalance < amount {
		return 0, fmt.Errorf("Invalid amount to transfer! You have %2.f and wants to transfer %2.f", accBalance/100, amt/100)
	}

	return accountOriginBalance, nil
}

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
