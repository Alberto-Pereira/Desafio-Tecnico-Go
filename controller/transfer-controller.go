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
// Create Transfer godoc
// @Summary Create a transfer
// @Description Needs to be logged before transfer. Returns a message associated with the operation
// @Tags Transfer
// @Accept application/json
// @Produce application/json
// @Param account body model.Transfer true "Only account_destination_id and amount(Note: Value in cents greater or equal to 1)"
// @Success 200 {string} message
// @Failure 400 {string} message
// @Router /transfers/ [post]
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
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, "Transfer created!")
}

// Read Transfers
// Receives an account id through a token from cookie and search for all transfers of that account
// If the operation is successful, returns one success code and the transfers of that account
// If the operation fails, returns one failure code and message
// Read Transfers godoc
// @Summary Read transfers from a logged account
// @Description Needs to be logged. Returns transfers
// @Tags Transfer
// @Accept application/json
// @Produce application/json
// @Success 200 {string} []model.Transfers
// @Failure 400 {string} message
// @Failure 404 {string} message
// @Router /transfers/ [get]
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
