package controller

import (
	"desafio-tecnico/model"
	"desafio-tecnico/service"
	"errors"

	"github.com/gin-gonic/gin"
)

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

func getTokenFromCookie(ctx *gin.Context) (string, error) {

	token, err := ctx.Cookie("token")
	if err != nil {
		return "", errors.New("Error while try to retrieve the token!")
	}

	return token, nil
}
