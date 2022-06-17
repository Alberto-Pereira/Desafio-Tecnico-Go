// Controller package contains controller operations
// for account, login and transfer models
// This file contains the routes for the API
package controller

import (
	"github.com/gin-gonic/gin"
)

// Setup Router
// Setups the router using custom groups of routes
// Returns the router
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
