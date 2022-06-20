package controller

import (
	"desafio-tecnico/security"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// Create Transfer
// Valid Transfer - New transfer with valid values
func TestCreateTransfer_ValidTransfer(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	token, expirationTime, _ := security.GenerateToken(91)

	cookie := http.Cookie{Name: "token", Value: token, MaxAge: expirationTime}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/transfers/", strings.NewReader(`{
		"account_destination_id": 93,
		"amount": 23
	}`))
	req.AddCookie(&cookie)

	router.ServeHTTP(w, req)

	assertions.Equal(w.Code, 200)
	assertions.Equal(w.Body.String(), "\"Transfer created!\"")
}

// Create Transfer
// Invalid Token - New transfer with invalid token
func TestCreateTransfer_InvalidToken(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	token, expirationTime, _ := security.GenerateToken(0)

	cookie := http.Cookie{Name: "", Value: token, MaxAge: expirationTime}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/transfers/", strings.NewReader(`{
		"account_destination_id": 93,
		"amount": 23
	}`))
	req.AddCookie(&cookie)

	router.ServeHTTP(w, req)

	assertions.Equal(w.Code, 400)
	assertions.Equal(w.Body.String(), "\"Error while try to retrieve the token!\"")
}

// Create Transfer
// Incorrect Transfer - New transfer with incorrect values
func TestCreateTransfer_IncorrectTransfer(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	token, expirationTime, _ := security.GenerateToken(91)

	cookie := http.Cookie{Name: "token", Value: token, MaxAge: expirationTime}

	w := httptest.NewRecorder()
	w2 := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/transfers/", strings.NewReader(`{
		"account_destination_id": "93",
		"amount": 23
	}`))
	req.AddCookie(&cookie)

	req2, _ := http.NewRequest("POST", "/transfers/", strings.NewReader(`{
		"account_destination_id": 93,
		"amount": "23"
	}`))
	req2.AddCookie(&cookie)

	router.ServeHTTP(w, req)
	router.ServeHTTP(w2, req2)

	assertions.Equal(w.Code, 400)
	assertions.Equal(w.Body.String(), "\"Error while try to retrieve transfer from request!\"")

	assertions.Equal(w2.Code, 400)
	assertions.Equal(w2.Body.String(), "\"Error while try to retrieve transfer from request!\"")
}

// Create Transfer
// Invalid Account Origin Id Transfer - New transfer with invalid account origin id
func TestCreateTransfer_InvalidAccountOriginIdTransfer(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	token, expirationTime, _ := security.GenerateToken(999)

	cookie := http.Cookie{Name: "token", Value: token, MaxAge: expirationTime}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/transfers/", strings.NewReader(`{
		"account_destination_id": 93,
		"amount": 23
	}`))
	req.AddCookie(&cookie)

	router.ServeHTTP(w, req)

	assertions.Equal(w.Code, 500)
	assertions.Equal(w.Body.String(), "\"Account origin not found!\"")
}

// Create Transfer
// Invalid Transfer - New transfer with invalid values
func TestCreateTransfer_InvalidTransfer(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	token, expirationTime, _ := security.GenerateToken(91)

	cookie := http.Cookie{Name: "token", Value: token, MaxAge: expirationTime}

	w := httptest.NewRecorder()
	w2 := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/transfers/", strings.NewReader(`{
		"account_destination_id": 999,
		"amount": 23
	}`))
	req.AddCookie(&cookie)

	req2, _ := http.NewRequest("POST", "/transfers/", strings.NewReader(`{
		"account_destination_id": 93,
		"amount": 0
	}`))
	req2.AddCookie(&cookie)

	router.ServeHTTP(w, req)
	router.ServeHTTP(w2, req2)

	assertions.Equal(w.Code, 400)
	assertions.Equal(w.Body.String(), "\"Account destination not found!\"")

	assertions.Equal(w2.Code, 400)
	assertions.Equal(w2.Body.String(), "\"Invalid amount to transfer!\"")
}

// Read Transfers
// Valid Account - Try to read transfers from a valid account
func TestReadTransfers_ValidAccount(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	token, expirationTime, _ := security.GenerateToken(91)

	cookie := http.Cookie{Name: "token", Value: token, MaxAge: expirationTime}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/transfers/", strings.NewReader(`{}`))
	req.AddCookie(&cookie)

	router.ServeHTTP(w, req)

	assertions.Equal(w.Code, 200)
	assertions.NotEmpty(w.Body.String())
}

// Read Transfers
// Invalid Token - Try to get a token to read the transfers
func TestReadTransfers_InvalidToken(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	token, expirationTime, _ := security.GenerateToken(0)

	cookie := http.Cookie{Name: "", Value: token, MaxAge: expirationTime}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/transfers/", strings.NewReader(`{}`))
	req.AddCookie(&cookie)

	router.ServeHTTP(w, req)

	assertions.Equal(w.Code, 400)
	assertions.Equal(w.Body.String(), "\"Error while try to retrieve the token!\"")
}

// Read Transfers
// Invalid Account - Try to read transfers from an invalid account
func TestReadTransfers_InvalidAccount(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	token, expirationTime, _ := security.GenerateToken(999)

	cookie := http.Cookie{Name: "token", Value: token, MaxAge: expirationTime}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/transfers/", strings.NewReader(`{}`))
	req.AddCookie(&cookie)

	router.ServeHTTP(w, req)

	assertions.Equal(w.Code, 404)
	assertions.Equal(w.Body.String(), "\"Account id not found!\"")
}

// Read Transfers
// Account Without Transfers - Try to read transfers from an account withour transfers
func TestReadTransfers_AccountWithoutTransfers(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	token, expirationTime, _ := security.GenerateToken(93)

	cookie := http.Cookie{Name: "token", Value: token, MaxAge: expirationTime}

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/transfers/", strings.NewReader(`{}`))
	req.AddCookie(&cookie)

	router.ServeHTTP(w, req)

	assertions.Equal(w.Code, 404)
	assertions.Equal(w.Body.String(), "\"This account doesn't have transfers!\"")
}
