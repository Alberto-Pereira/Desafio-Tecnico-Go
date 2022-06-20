package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// Create Account
// Valid Account - New account with valid values
func TestCreateAccount_ValidAccount(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	// ! Cpf needs to be change before every test after the first test
	req, _ := http.NewRequest("POST", "/accounts/", strings.NewReader(`{
		"name": "Test User",
		"cpf": "119.112.111-11",
		"secret": "123456",
		"balance": 1000
	}`))

	router.ServeHTTP(w, req)

	assertions.Equal(w.Code, 200)
	assertions.Equal(w.Body.String(), "\"Account created!\"")
}

// Create Account
// Incorrect Account - New account with incorrect fields
func TestCreateAccount_IncorrectAccount(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()
	w2 := httptest.NewRecorder()
	w3 := httptest.NewRecorder()
	w4 := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/accounts/", strings.NewReader(`{
		"name": Test User,
		"cpf": "119.112.111-11",
		"secret": "123456",
		"balance": 1000
	}`))

	req2, _ := http.NewRequest("POST", "/accounts/", strings.NewReader(`{
		"name": "Test User",
		"cpf": 119.112.111-11,
		"secret": "123456",
		"balance": 1000
	}`))

	req3, _ := http.NewRequest("POST", "/accounts/", strings.NewReader(`{
		"name": "Test User",
		"cpf": "119.112.111-11",
		"secret": 123456,
		"balance": 1000
	}`))

	req4, _ := http.NewRequest("POST", "/accounts/", strings.NewReader(`{
		"name": "Test User",
		"cpf": "119.112.111-11",
		"secret": "123456",
		"balance": "1000"
	}`))

	router.ServeHTTP(w, req)
	router.ServeHTTP(w2, req2)
	router.ServeHTTP(w3, req3)
	router.ServeHTTP(w4, req4)

	assertions.Equal(w.Code, 400)
	assertions.Equal(w.Body.String(), "\"Error while try to retrieve data from request!\"")

	assertions.Equal(w2.Code, 400)
	assertions.Equal(w2.Body.String(), "\"Error while try to retrieve data from request!\"")

	assertions.Equal(w3.Code, 400)
	assertions.Equal(w3.Body.String(), "\"Error while try to retrieve data from request!\"")

	assertions.Equal(w4.Code, 400)
	assertions.Equal(w4.Body.String(), "\"Error while try to retrieve data from request!\"")
}

// Create Account
// Invalid Account Values - New account with invalid values
func TestCreateAccount_InvalidAccountValues(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()
	w2 := httptest.NewRecorder()
	w3 := httptest.NewRecorder()
	w4 := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/accounts/", strings.NewReader(`{
		"name": "test User",
		"cpf": "119.112.111-92",
		"secret": "123456",
		"balance": 1000
	}`))

	req2, _ := http.NewRequest("POST", "/accounts/", strings.NewReader(`{
		"name": "Test User",
		"cpf": "1219.112.111-92",
		"secret": "123456",
		"balance": 1000
	}`))

	req3, _ := http.NewRequest("POST", "/accounts/", strings.NewReader(`{
		"name": "Test User",
		"cpf": "119.112.111-92",
		"secret": " ",
		"balance": 1000
	}`))

	req4, _ := http.NewRequest("POST", "/accounts/", strings.NewReader(`{
		"name": "Test User",
		"cpf": "119.112.111-92",
		"secret": "123456",
		"balance": -1
	}`))

	router.ServeHTTP(w, req)
	router.ServeHTTP(w2, req2)
	router.ServeHTTP(w3, req3)
	router.ServeHTTP(w4, req4)

	assertions.Equal(w.Code, 400)
	assertions.Equal(w.Body.String(), "\"Error validating account!\"")

	assertions.Equal(w2.Code, 400)
	assertions.Equal(w2.Body.String(), "\"Error validating account!\"")

	assertions.Equal(w3.Code, 400)
	assertions.Equal(w3.Body.String(), "\"Error validating account!\"")

	assertions.Equal(w4.Code, 400)
	assertions.Equal(w4.Body.String(), "\"Error validating account!\"")
}

// Create Account
// Invalid Account - New account with cpf already registred
func TestCreateAccount_InvalidAccount(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/accounts/", strings.NewReader(`{
		"name": "Test User",
		"cpf": "119.112.111-11",
		"secret": "123456",
		"balance": 1000
	}`))

	router.ServeHTTP(w, req)

	assertions.Equal(w.Code, 400)
	assertions.Equal(w.Body.String(), "\"Account already created!\"")
}

// Read Account Balance
// Valid Account - Try to read from a valid account balance
func TestReadAccountBalance_ValidAccount(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/accounts/66/balance", strings.NewReader(`{}`))

	router.ServeHTTP(w, req)

	assertions.Equal(w.Code, 200)
	assertions.Equal(w.Body.String(), "100")
}

// Read Account Balance
// Incorrect Account - Try to read from an incorrect account balance
func TestReadAccountBalance_IncorrectAccount(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()
	w2 := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/accounts/0/balance", strings.NewReader(`{}`))

	req2, _ := http.NewRequest("GET", "/accounts/invalidId/balance", strings.NewReader(`{}`))

	router.ServeHTTP(w, req)
	router.ServeHTTP(w2, req2)

	assertions.Equal(w.Code, 400)
	assertions.Equal(w.Body.String(), "\"Invalid account id!\"")

	assertions.Equal(w2.Code, 400)
	assertions.Equal(w2.Body.String(), "\"Invalid account id!\"")
}

// Read Account Balance
// Invalid Account - Try to read from an invalid account balance
func TestReadAccountBalance_InvalidAccount(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/accounts/1/balance", strings.NewReader(`{}`))

	router.ServeHTTP(w, req)

	assertions.Equal(w.Code, 404)
	assertions.Equal(w.Body.String(), "\"This account doesn't exist!\"")
}

// Read Accounts
// Accounts Registred - Existing accounts
func TestReadAccounts_AccountsRegistred(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/accounts/", strings.NewReader(`{}`))

	router.ServeHTTP(w, req)

	assertions.Equal(w.Code, 200)
	assertions.NotEmpty(w.Body.String())
}

// Read Accounts
// Accounts Not Registred - Not existing accounts
func TestReadAccounts_AccountsNotRegistred(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	// Needs to delete all registred accounts before testing
	req, _ := http.NewRequest("GET", "/accounts/", strings.NewReader(`{}`))

	router.ServeHTTP(w, req)

	assertions.Equal(w.Code, 404)
	assertions.Equal(w.Body.String(), "\"No accounts registred!\"")
}
