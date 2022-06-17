// Model package contains the entities of the application
package model

// Login is a simplification of account model,
// containing only the cpf and secret
type Login struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}
