package models

import (
	"fmt"
	"net/http"
)

//Account structure, used to manipulate accounts
type Account struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Cpf        string  `json:"cpf"`
	Secret     string  `json:"secret"`
	Balance    float64 `json:"balance"`
	Created_at string  `json:"created_at"`
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
