package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// Login
// Valid Login - Valid login values
func TestLogin_ValidLogin(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/login", strings.NewReader(`{
		"cpf": "119.112.111-11",
		"secret": "123456"
	}`))

	router.ServeHTTP(w, req)

	assertions.Equal(w.Code, 200)
	assertions.Equal(w.Body.String(), "\"Logged!\"")
}

// Login
// Incorrect Login - Incorrect login values
func TestLogin_IncorrectLogin(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()
	w2 := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/login", strings.NewReader(`{
		"cpf": 119.112.111-11,
		"secret": "123456"
	}`))

	req2, _ := http.NewRequest("POST", "/login", strings.NewReader(`{
		"cpf": "119.112.111-11",
		"secret": 123456
	}`))

	router.ServeHTTP(w, req)
	router.ServeHTTP(w2, req2)

	assertions.Equal(w.Code, 400)
	assertions.Equal(w.Body.String(), "\"Error while try to retrieve data from login!\"")

	assertions.Equal(w.Code, 400)
	assertions.Equal(w.Body.String(), "\"Error while try to retrieve data from login!\"")
}

// Login
// Invalid Login - Invalid login values
func TestLogin_InvalidLogin(t *testing.T) {

	router := SetupRouter()

	assertions := require.New(t)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/login", strings.NewReader(`{
		"cpf": "129.112.111-11",
		"secret": "123456"
	}`))

	router.ServeHTTP(w, req)

	assertions.Equal(w.Code, 404)
	assertions.Equal(w.Body.String(), "\"Invalid account cpf!\"")
}
