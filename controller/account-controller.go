// Controller package contains controller operations
// for account, login and transfer models
package controller

import (
	"desafio-tecnico/model"
	"desafio-tecnico/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Account
// Receives an account through a request, bind in one account model and send to the service
// If the operation is successful, returns one success code and message
// If the operation fails, returns one failure code and message
// Create Account godoc
// @Summary Create an account
// @Description Returns a message associated with the operation
// @Tags Account
// @Accept application/json
// @Produce application/json
// @Param account body model.Account true "Only name(Ex.: Name), cpf(Ex.: 000.000.000-00), secret(Note: Anything but whitespace) and balance(Note: Value in cents greater or equal to 0) "
// @Success 200 {string} message
// @Failure 400 {string} message
// @Router /accounts/ [post]
func CreateAccount(ctx *gin.Context) {

	var account model.Account

	err := ctx.BindJSON(&account)
	if err != nil {
		ctx.JSON(400, "Error while try to retrieve data from request!")
		return
	}

	err = service.CreateAccount(account)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, "Account created!")
}

// Read Account Balance
// Receives an account id through a request, bind in one variable and send to the service
// If the operation is successful, returns one success code and the account balance
// If the operation fails, returns one failure code and message
// Read Account Balance godoc
// @Summary Read a balance from an account
// @Description Returns an account balance
// @Tags Account
// @Accept application/json
// @Produce application/json
// @Param account_id path int true "Inform the account id"
// @Success 200 {string} message
// @Failure 400 {string} message
// @Failure 404 {string} message
// @Router /accounts/{account_id}/balance [get]
func ReadAccountBalance(ctx *gin.Context) {

	accountId := ctx.Param("account_id")

	accId, err := strconv.Atoi(accountId)
	if err != nil || accId <= 0 {
		ctx.JSON(400, "Invalid account id!")
		return
	}

	accountBalance, err := service.ReadAccountBalance(accId)
	if err != nil {
		ctx.JSON(404, err.Error())
		return
	}

	accBalance := float64(accountBalance)

	ctx.JSON(200, accBalance/100)
}

// Read Accounts
// Search for all accounts created in the database
// If the operation is successful, returns one success code and the accounts
// If the operation fails, returns one failure code and message
// Read Accounts godoc
// @Summary Read accounts registred in database
// @Description Returns accounts
// @Tags Account
// @Accept application/json
// @Produce application/json
// @Success 200 {string} []model.Account
// @Failure 404 {string} message
// @Router /accounts/ [get]
func ReadAccounts(ctx *gin.Context) {

	accounts, err := service.ReadAccounts()
	if err != nil {
		ctx.JSON(404, err.Error())
		return
	}

	ctx.JSON(200, accounts)
}
