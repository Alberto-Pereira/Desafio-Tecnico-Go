// Model package contains the entities of the application
package model

type Account struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	CPF        string `json:"cpf"`
	Secret     string `json:"secret"`
	Balance    int    `json:"balance"`
	Created_at int    `json:"created_at"`
}
