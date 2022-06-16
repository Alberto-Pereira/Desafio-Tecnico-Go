package controller

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	accounts := router.Group("/accounts")
	{
		accounts.GET("/", ReadAccounts)

		accounts.GET("/:account_id/balance", ReadAccountBalance)

		accounts.POST("/", CreateAccount)
	}

	router.POST("/login", Login)

	transfers := router.Group("/transfers")
	{
		transfers.GET("/", ReadTransfers)
		transfers.POST("/", CreateTransfer)
	}

	return router
}
