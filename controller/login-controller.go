package controller

import (
	"desafio-tecnico/model"
	"desafio-tecnico/security"
	"desafio-tecnico/service"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {

	var login model.Login

	err := ctx.BindJSON(&login)

	if err != nil {
		ctx.JSON(400, "Bad request!")
		return
	}

	accountId, err := service.ReadAccount(login)

	if err != nil {
		ctx.JSON(404, err.Error())
		return
	}

	token, expirationTime, err := security.GenerateToken(accountId)

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.SetCookie("token", token, expirationTime, "/", "localhost", true, true)

	ctx.JSON(200, "Logged!")
}
