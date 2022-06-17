// Controller package contains controller operations
// for account, login and transfer models
package controller

import (
	"desafio-tecnico/model"
	"desafio-tecnico/security"
	"desafio-tecnico/service"

	"github.com/gin-gonic/gin"
)

// Login
// Receives a login through a request, bind in one login model,
// search for an account that matches the login and generate a token
// If the operation is successful, set the token in the cookie then returns one success code and message
// If the operation fails, returns one failure code and message
func Login(ctx *gin.Context) {

	var login model.Login

	err := ctx.BindJSON(&login)
	if err != nil {
		ctx.JSON(400, "Error while try to retrieve data from login!")
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
