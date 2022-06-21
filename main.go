package main

import (
	"desafio-tecnico/controller"
	"desafio-tecnico/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Desafio Técnico GO
// @version         1.0
// @description     API for transfer between internal accounts of a digital bank

// @contact.name   Desafio Técnico GO Support
// @contact.url    https://github.com/Alberto-Pereira/desafio-tecnico-go
// @contact.email  alberto.pereira.dev@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
func main() {

	// Receives the router
	router := controller.SetupRouter()

	// Swagger path
	docs.SwaggerInfo.BasePath = "/"

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Start the router
	router.Run()
}
