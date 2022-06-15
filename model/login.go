package model

type Login struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}
