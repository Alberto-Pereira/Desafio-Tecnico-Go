// Service package contains service operations
// for account and transfer models
package service

import (
	"desafio-tecnico/model"
	"desafio-tecnico/repository"
	"errors"
	"fmt"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Create Account
// Receives an account, validates and send to the repository
// If the operation is successful, returns nil
// If the operation fails, returns an error
func CreateAccount(account model.Account) error {

	err := validateAccount(account)
	if err != nil {
		return fmt.Errorf("Error validating account! %s", err.Error())
	}

	err = checkAccountCpf(account.CPF)
	if err != nil {
		return err
	}

	hashSecret, err := hashAccountSecret(account.Secret)
	if err != nil {
		return err
	}

	err = repository.CreateAccount(model.Account{
		Name: account.Name, CPF: account.CPF, Secret: hashSecret,
		Balance: account.Balance, Created_at: int(time.Now().Unix())})
	if err != nil {
		return err
	}

	return nil
}

// Read Account
// Receives an login, validates and search for the account id
// If the operation is successful, returns the account id and nil
// If the operation fails, returns 0 and an error
func ReadAccount(login model.Login) (int, error) {

	err := validateCpfAndSecret(login)
	if err != nil {
		return 0, err
	}

	accountId, accountSecretHash, err := repository.ReadAccount(login.CPF)
	if err != nil {
		return 0, err
	}

	err = checkAccountSecretHash(login.Secret, accountSecretHash)
	if err != nil {
		return 0, errors.New("Incorrect account secret!")
	}

	return accountId, nil
}

// Read Account Balance
// Receives an account id and search for the account balance
// If the operation is successful, returns the account balance and nil
// If the operation fails, returns 0 and an error
func ReadAccountBalance(accountId int) (int, error) {

	accountBalance, err := repository.ReadAccountBalance(accountId)
	if err != nil {
		return 0, err
	}

	return accountBalance, nil
}

// Read Accounts
// Search in the repository for all accounts created
// If the operation is successful, returns the accounts and nil
// If the operation fails, returns nil and an error
func ReadAccounts() ([]model.Account, error) {

	accounts, err := repository.ReadAccounts()
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

// Validate Account
// Receives an account and validates using custom formats
// If the operation is successful, returns nil
// If the operation fails, returns an error
func validateAccount(account model.Account) error {

	// ! Format - First letter uppercase, at least 2 characters and accepts space
	validName := regexp.MustCompile(`^([A-Z]{1}[a-z]+\s?)+$`)
	// ! Format - Accepts the following 000.000.000-00 pattern
	validCPF := regexp.MustCompile("^[0-9]{3}[.][0-9]{3}[.][0-9]{3}[-][0-9]{2}$")
	// ! Format - Matches anything other than a space, tab or newline
	validSecret := regexp.MustCompile(`^\S+$`)

	if !validName.MatchString(account.Name) {
		return errors.New("Invalid account name!")
	}

	if !validCPF.MatchString(account.CPF) {
		return errors.New("Invalid account cpf!")
	}

	if !validSecret.MatchString(account.Secret) {
		return errors.New("Invalid account secret!")
	}

	if account.Balance < 0 {
		return errors.New("Invalid account balance!")
	}

	return nil
}

// Check Account Cpf
// Receives an account cpf and search for that cpf
// If there's no cpf, returns nil
// If the operation fails or the cpf exists, returns an error
func checkAccountCpf(accountCpf string) error {

	isAccountCreated, err := repository.ReadAccountCpf(accountCpf)
	if err != nil {
		return err
	}

	if isAccountCreated {
		return errors.New("Account already created!")
	}

	return nil
}

// Hash Account Secret
// Receives an account secret and hashes that secret
// If the operation is successful, returns that secret in hash format and nil
// If the operation fails, returns "" and an error
func hashAccountSecret(accountSecret string) (string, error) {

	hashSecret, err := bcrypt.GenerateFromPassword([]byte(accountSecret), 14)
	if err != nil {
		return "", errors.New("Error while try to hash account secret!")
	}

	err = checkAccountSecretHash(accountSecret, string(hashSecret))
	if err != nil {
		return "", errors.New("Error while try to check if account secret and account secret hash matches!")
	}

	return string(hashSecret), nil
}

// Check Account Secret Hash
// Receives an account secret and hash secret and verify if the secrets matches
// If the operation is successful, returns nil
// If the operation fails, returns an error
func checkAccountSecretHash(accountSecret string, hashSecret string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hashSecret), []byte(accountSecret))

	return err
}

// Validate Cpf And Secret
// Receives a login and validates using custom formats
// If the operation is successful, returns nil
// If the operation fails, returns an error
func validateCpfAndSecret(login model.Login) error {

	// ! Format - Accepts the following 000.000.000-00 pattern
	validCPF := regexp.MustCompile("^[0-9]{3}[.][0-9]{3}[.][0-9]{3}[-][0-9]{2}$")
	// ! Format - Matches anything other than a space, tab or newline
	validSecret := regexp.MustCompile(`^\S+$`)

	if !validCPF.MatchString(login.CPF) {
		return errors.New("Invalid account cpf!")
	}

	if !validSecret.MatchString(login.Secret) {
		return errors.New("Invalid account secret!")
	}

	return nil
}
