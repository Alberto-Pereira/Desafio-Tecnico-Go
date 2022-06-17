// Security package contains the files for application security
package security

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("desafio-tecnico-go")

// Custom standard claims with account id
type Claims struct {
	AccountID int `json:"account_origin_id"`
	jwt.StandardClaims
}

// Generate Token
// Receives an account id and generates a token
// If the operation is successful, returns the signed token, an expiration time and nil
// If the operation fails, returns "", 0 and an error
func GenerateToken(accountId int) (string, int, error) {

	expirationTime := time.Now().Add(time.Minute * 5).Unix()

	claims := Claims{
		AccountID: accountId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", 0, errors.New("Error while try to create token!")
	}

	return signedToken, int(expirationTime), nil
}

// Validate Token
// Receives a token and validates
// If the operation is successful, returns an account id and nil
// If the operation fails, returns 0 and an error
func ValidateToken(token string) (int, error) {

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !tkn.Valid {
		return 0, errors.New("Invalid token!")
	}

	return claims.AccountID, nil
}
