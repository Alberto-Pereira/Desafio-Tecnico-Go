package main

import "desafio-tecnico/controller"

func main() {

	router := controller.SetupRouter()

	router.Run()
}
