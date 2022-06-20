package service

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
		Name: "Test User", CPF: "777.122.122-00", Secret: "TEST!TEST!TEST!", Balance: 1, Created_at: 1655361423}

	err := CreateAccount(validAccount)

	assertions.Empty(err)
}

// Create Account
// Invalid Account - New account with invalid values
func TestCreateAccount_InvalidAccount(t *testing.T) {

	assertions := require.New(t)

	invalidAccount := []model.Account{
		{Name: "test User", CPF: "122.122.122-00", Secret: "TEST!TEST!TEST!", Balance: 10000, Created_at: 1655361423},
		{Name: "Test User", CPF: "1122.122.122-00", Secret: "TEST!TEST!TEST!", Balance: 10000, Created_at: 1655361423},
		{Name: "Test User", CPF: "122.122.122-00", Secret: " ", Balance: 10000, Created_at: 1655361423},
		{Name: "Test User", CPF: "122.122.122-00", Secret: "TEST!TEST!TEST!", Balance: -1, Created_at: 1655361423},
		{},
	}

	for _, invalidAcc := range invalidAccount {
		err := CreateAccount(invalidAcc)
		assertions.NotEmpty(err)
	}
}

// Create Account
// Account Already Created - New account with valid values but cpf already registred
func TestCreateAccount_AccountAlreadyCreated(t *testing.T) {

	assertions := require.New(t)

	validAccount := model.Account{
		Name: "Test User", CPF: "122.122.122-00", Secret: "TEST!TEST!TEST!", Balance: 10000, Created_at: 1655361423}

	err := CreateAccount(validAccount)

	assertions.Equal(err.Error(), "Account already created!")
}

// Read Account
// Valid Login - Login with valid values
func TestReadAccount_ValidLogin(t *testing.T) {

	assertions := require.New(t)

	validLogin := model.Login{CPF: "122.122.122-00", Secret: "TEST!TEST!TEST!"}

	accountId, err := ReadAccount(validLogin)

	assertions.Equal(accountId, 61)
	assertions.Empty(err)
}

// Read Account
// Invalid Login - Login with invalid credentials
func TestReadAccount_InvalidLogin(t *testing.T) {

	assertions := require.New(t)

	// ! Cpf needs to be change before every test after the first test
	validLogin := []model.Login{
		{CPF: "1122.122.122-00", Secret: "TEST!TEST!TEST!"},
		{CPF: "122.122.122-00", Secret: " "},
		{},
	}

	for _, validLgn := range validLogin {
		accountId, err := ReadAccount(validLgn)

		assertions.Equal(accountId, 0)
		assertions.NotEmpty(err)
	}
}

// Read Account
// Login Invalid Cpf - Login with invalid cpf
func TestReadAccount_LoginInvalidCpf(t *testing.T) {

	assertions := require.New(t)

	loginInvalidCpf := model.Login{CPF: "999.999.999-99", Secret: "TEST!TEST!TEST!"}

	accountId, err := ReadAccount(loginInvalidCpf)

	assertions.Equal(accountId, 0)
	assertions.Equal(err.Error(), "Invalid account cpf!")
}

// Read Account
// Login Invalid Secret - Login with invalid secret
func TestReadAccount_LoginInvalidSecret(t *testing.T) {

	assertions := require.New(t)

	loginInvalidSecret := model.Login{CPF: "122.122.122-00", Secret: "INVALIDSECRET"}

	accountId, err := ReadAccount(loginInvalidSecret)

	assertions.Equal(accountId, 0)
	assertions.Equal(err.Error(), "Incorrect account secret!")
}

// Read Account Balance
// Valid Account - Account with valid id
func TestReadAccountBalance_ValidAccount(t *testing.T) {

	assertions := require.New(t)

	validAccount := model.Account{ID: 61}

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
// Accounts Created - Accounts registred in database
func TestReadAccounts_AccountsCreated(t *testing.T) {

	assertions := require.New(t)

	accounts, err := ReadAccounts()

	assertions.NotEmpty(accounts)
	assertions.Empty(err)
}

// Read Accounts
// Accounts Not Created - Accounts not registred in database
func TestReadAccounts_AccountsNotCreated(t *testing.T) {

	assertions := require.New(t)

	// Needs to delete all registred accounts before testing
	accounts, err := ReadAccounts()

	assertions.Empty(accounts)
	assertions.Equal(err.Error(), "No accounts registred!")
}
