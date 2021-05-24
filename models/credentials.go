package models

import (
	"fmt"
	"net/http"
)

//Credentials structure for manipulation
type Credentials struct {
	Cpf        string `json:"cpf"`
	Secret     string `json:"secret"`
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
