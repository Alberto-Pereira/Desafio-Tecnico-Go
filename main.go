package main

import (
	"desafio-tecnico/service"
	"fmt"
)

func main() {

	// id, err := service.CreateAccount(model.Account{
	// 	Name:       "Alberto",
	// 	CPF:        "700.123.123-00",
	// 	Secret:     "123456",
	// 	Balance:    10000,
	// 	Created_at: 1655239680,
	// })

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(id)

	accs, err := service.ReadAccounts()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(accs)
}
