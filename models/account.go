package models

import (
	"fmt"
	"net/http"
)

type Account struct {
	Id        int `json:"id"`
	Name      string
	Cpf       string
	Secret    string
	Balance   float64
	CreatedAt string `json:"created_at"`
}

type AccountList struct {
	Accounts []Account `json:"accounts"`
}

func (a *Account) Bind(r *http.Request) error {
	if a.Name == "" {
		return fmt.Errorf("name is a required field")
	}
	return nil
}
func (*AccountList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (*Account) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
