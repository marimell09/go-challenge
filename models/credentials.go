package models

import (
	"fmt"
	"net/http"
)

//Credentials structure for manipulation
type Credentials struct {
	// the cpf for login
	//
	// required: true
	// example: 99999999999
	Cpf string `json:"cpf"`

	// the secret for login
	//
	// requires: true
	// example: senha
	Secret string `json:"secret"`

	// the account id related to the cpf
	//
	// read-only: true
	// example: 1
	Account_id string `json:"account_id"`
}

//Bind method for credentials manipulation
func (c *Credentials) Bind(r *http.Request) error {
	if c.Cpf == "" || c.Secret == "" {
		return fmt.Errorf("cpf and secret are required fields")
	}
	return nil
}

//Render method for credentials
func (*Credentials) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
