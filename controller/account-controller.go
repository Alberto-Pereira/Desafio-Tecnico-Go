package controller

import (
	"desafio-tecnico/model"
	"desafio-tecnico/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ReadAccounts(ctx *gin.Context) {

	accounts, err := service.ReadAccounts()

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, accounts)
}

func ReadAccountBalance(ctx *gin.Context) {

	accountId := ctx.Param("account_id")

	accId, err := strconv.Atoi(accountId)

	if err != nil || accId <= 0 {
		ctx.JSON(500, "Invalid account id!")
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

func CreateAccount(ctx *gin.Context) {

	var account model.Account

	err := ctx.BindJSON(&account)

	if err != nil {
		ctx.JSON(400, "Error while try to retrieve data from request!")
		return
	}

	err = service.CreateAccount(account)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, "Account created!")
}
