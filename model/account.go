// Model package contains the entities of the application
package model

type Account struct {
	ID         int64
	Name       string
	CPF        string
	Secret     string
	Balance    int64
	Created_at int64
}
