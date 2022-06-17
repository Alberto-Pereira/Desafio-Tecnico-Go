// Controller package contains controller operations
// for account, login and transfer models
package controller

import (
	"desafio-tecnico/model"
	"desafio-tecnico/service"
	"errors"

	"github.com/gin-gonic/gin"
)

// Create Transfer
// Receives a transfer through a request, get token from cookie, bind in one transfer model
// then send to the service
// If the operation is successful, returns one success code and message
// If the operation fails, returns one failure code and message
func CreateTransfer(ctx *gin.Context) {

	token, err := getTokenFromCookie(ctx)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	var transfer model.Transfer

	err = ctx.BindJSON(&transfer)
	if err != nil {
		ctx.JSON(400, "Error while try to retrieve transfer from request!")
		return
	}

	err = service.CreateTransfer(token, transfer.Account_destination_id, transfer.Amount)
	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, "Transfer created!")
}

// Read Transfers
// Receives an account id through a token from cookie and search for all transfers of that account
// If the operation is successful, returns one success code and the transfers of that account
// If the operation fails, returns one failure code and message
func ReadTransfers(ctx *gin.Context) {

	token, err := getTokenFromCookie(ctx)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	transfers, err := service.ReadTransfers(token)
	if err != nil {
		ctx.JSON(404, err.Error())
		return
	}

	ctx.JSON(200, transfers)
}

// Get Token From Cookie
// Receives a request and search for a token in cookie
// If the operation is successful, returns the token and nil
// If the operation fails, returns "" and an error
func getTokenFromCookie(ctx *gin.Context) (string, error) {

	token, err := ctx.Cookie("token")
	if err != nil {
		return "", errors.New("Error while try to retrieve the token!")
	}

	return token, nil
}
