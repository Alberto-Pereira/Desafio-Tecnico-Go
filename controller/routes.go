package controller

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	accounts := router.Group("/accounts")
	{
		accounts.GET("/")

		accounts.GET("/:account_id/balance")

		accounts.POST("/")
	}

	return router
}
