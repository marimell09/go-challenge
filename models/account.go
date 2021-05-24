package models

import (
	"fmt"
	"net/http"
)

//Account structure, used to manipulate accounts
// swagger:model
type Account struct {
	// the id for this account
	//
	// read-only: true
	// in:query
	// example: 1
	Id int `json:"id"`

	// the user name
	//
	// required: true
	// example: Mariana
	Name string `json:"name"`

	// the user cpf
	//
	// required: true
	// example: 99999999999
	Cpf string `json:"cpf"`

	// the user secret
	//
	// required: true
	// example: senha
	Secret string `json:"secret"`

	// the account balance
	//
	// required: true
	// example: 150.50
	Balance float64 `json:"balance"`

	// the account creation date
	//
	// read-only: true
	// example: 2021-05-24T03:30:26.021292Z
	Created_at string `json:"created_at"`
}

//Account list
type AccountList struct {
	Accounts []Account `json:"accounts"`
}

//Bind method to account manipulation
func (a *Account) Bind(r *http.Request) error {
	if a.Name == "" || a.Cpf == "" || a.Secret == "" {
		return fmt.Errorf("missing required fields")
	}
	return nil
}

//Render method for account list
func (*AccountList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

//Render method for account
func (*Account) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
