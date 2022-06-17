package security

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("desafio-tecnico-go")

type Claims struct {
	AccountID int `json:"account_origin_id"`
	jwt.StandardClaims
}

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
