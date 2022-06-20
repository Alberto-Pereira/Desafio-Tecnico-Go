package repository

import (
	"desafio-tecnico/model"
	"testing"

	"github.com/stretchr/testify/require"
)

// Create Account
// Valid Account - New account with valid values
func TestCreateAccount_ValidAccount(t *testing.T) {

	assertions := require.New(t)

	// ! Cpf needs to be change before every test after the first test
	validAccount := model.Account{
		Name: "Test User", CPF: "773.123.123-00", Secret: "TEST!TEST!TEST!", Balance: 1, Created_at: 1655361423}

	err := CreateAccount(validAccount)

	assertions.Empty(err)
}

// Create Account
// Invalid Account - New account with invalid fields
func TestCreateAccount_InvalidAccount(t *testing.T) {

	assertions := require.New(t)

	invalidAccount := []model.Account{
		{Name: "Test User", CPF: "123.123.123-00", Secret: "TEST!TEST!TEST!", Balance: 10000, Created_at: 1655361423},
		{CPF: "023.023.023-00", Secret: "TEST!TEST!TEST!", Balance: 10000, Created_at: 1655361423},
		{Name: "Test User", Secret: "TEST!TEST!TEST!", Balance: 10000, Created_at: 1655361423},
		{Name: "Test User", CPF: "023.023.023-00", Balance: 10000, Created_at: 1655361423},
		{Name: "Test User", CPF: "023.023.023-00", Secret: "TEST!TEST!TEST!", Balance: 10000},
		{},
	}

	for _, invalidAcc := range invalidAccount {
		err := CreateAccount(invalidAcc)

		assertions.Equal(err.Error(), "Error while try to create account!")
	}
}

// Read Account
// Valid Account - Account with valid cpf
func TestReadAccount_ValidAccount(t *testing.T) {

	assertions := require.New(t)

	validAccount := model.Account{CPF: "123.123.123-00"}

	accountId, accountSecret, err := ReadAccount(validAccount.CPF)

	assertions.Equal(accountId, 35)
	assertions.Equal(accountSecret, "TEST!TEST!TEST!")
	assertions.Empty(err)
}

// Read Account
// Invalid Account - Account with invalid cpf
func TestReadAccount_InvalidAccount(t *testing.T) {

	assertions := require.New(t)

	invalidAccount := model.Account{CPF: "000.000.000-00"}

	accountId, accountSecret, err := ReadAccount(invalidAccount.CPF)

	assertions.Equal(accountId, 0)
	assertions.Equal(accountSecret, "")
	assertions.Equal(err.Error(), "Invalid account cpf!")
}

// Read Account Id
// Valid Account - Account with valid id
func TestReadAccountId_ValidAccount(t *testing.T) {

	assertions := require.New(t)

	validAccount := model.Account{ID: 35}

	err := ReadAccountId(validAccount.ID)

	assertions.Empty(err)
}

// Read Account Id
// Invalid Account - Account with invalid id
func TestReadAccountId_InvalidAccount(t *testing.T) {

	assertions := require.New(t)

	invalidAccount := model.Account{ID: 0}

	err := ReadAccountId(invalidAccount.ID)

	assertions.Equal(err.Error(), "Account id not found!")
}

// Read Account Cpf
// Valid Account - Account with valid cpf
func TestReadAccountCpf_ValidAccount(t *testing.T) {

	assertions := require.New(t)

	validAccount := model.Account{CPF: "123.123.123-00"}

	err := ReadAccountCpf(validAccount.CPF)

	assertions.Empty(err)
}

// Read Account Cpf
// Valid Account With No Registred Cpf - Account with no registred cpf
func TestReadAccountCpf_ValidAccountWithNoRegistredCpf(t *testing.T) {

	assertions := require.New(t)

	validAccount := model.Account{CPF: "000.000.000-00"}

	err := ReadAccountCpf(validAccount.CPF)

	assertions.Equal(err.Error(), "Error while try to search for account cpf!")
}

// Read Account Balance
// Valid Account - Account with valid id
func TestReadAccountBalance_ValidAccount(t *testing.T) {

	assertions := require.New(t)

	validAccount := model.Account{ID: 35}

	accountBalance, err := ReadAccountBalance(validAccount.ID)

	assertions.Equal(accountBalance, 10000)
	assertions.Empty(err)
}

// Read Account Balance
// Invalid Account - Account with invalid id
func TestReadAccountBalance_InvalidAccount(t *testing.T) {

	assertions := require.New(t)

	invalidAccount := model.Account{ID: 0}

	accountBalance, err := ReadAccountBalance(invalidAccount.ID)

	assertions.Equal(accountBalance, 0)
	assertions.Equal(err.Error(), "This account doesn't exist!")
}

// Read Accounts
// Existing Accounts - Existing accounts
func TestReadAccounts_ExistingAccounts(t *testing.T) {

	assertions := require.New(t)

	accounts, err := ReadAccounts()

	assertions.NotEmpty(accounts)
	assertions.Empty(err)
}

// Read Accounts
// Non Existing Accounts - Non existing accounts
func TestReadAccounts_NonExistingAccounts(t *testing.T) {

	assertions := require.New(t)

	// Needs to delete all registred accounts before testing
	accounts, err := ReadAccounts()

	assertions.Empty(accounts)
	assertions.Equal(err.Error(), "Error while try to read accounts!")
}

// Update Account Balance
// Valid Account - Account with valid id
func TestUpdateAccountBalance_ValidAccount(t *testing.T) {

	assertions := require.New(t)

	validAccount := model.Account{ID: 35, Balance: 20000}

	err := UpdateAccountBalance(validAccount.ID, validAccount.Balance)

	assertions.Empty(err)
}

// Update Account Balance
// Invalid Account - Account with invalid id
func TestUpdateAccountBalance_InvalidAccount(t *testing.T) {

	assertions := require.New(t)

	invalidAccount := model.Account{ID: 0, Balance: 20000}

	err := UpdateAccountBalance(invalidAccount.ID, invalidAccount.Balance)

	assertions.Equal(err.Error(), "Error while try to update account balance!")
}
