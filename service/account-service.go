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

func ReadAccounts() ([]model.Account, error) {

	accounts, err := repository.ReadAccounts()
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func ReadAccountBalance(accountId int) (int, error) {

	accountBalance, err := repository.ReadAccountBalance(accountId)
	if err != nil {
		return 0, err
	}

	return accountBalance, nil
}

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

func checkAccountSecretHash(accountSecret string, hashSecret string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hashSecret), []byte(accountSecret))

	return err
}

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
