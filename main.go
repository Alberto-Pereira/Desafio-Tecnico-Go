package main

import "desafio-tecnico/controller"

func main() {

	// Receives the router
	router := controller.SetupRouter()

	// Start the router
	router.Run()
}
