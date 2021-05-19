package models

import (
	"fmt"
	"net/http"
)

type Credentials struct {
	Cpf    string `json:"cpf"`
	Secret string `json:"secret"`
}

func (c *Credentials) Bind(r *http.Request) error {
	if c.Cpf == "" || c.Secret == "" {
		return fmt.Errorf("cpf and secret are required fields")
	}
	return nil
}

func (*Credentials) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
