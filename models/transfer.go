package models

import (
	"fmt"
	"net/http"
)

//Transfer structure for transfer manipulations
type Transfer struct {
	// the transfer id
	//
	// read-only: true
	// example: 1
	Id int `json:"id"`

	// the trasnfer origin id
	//
	// read-only: true
	// example: 1
	Account_origin_id int `json:"account_origin_id"`

	// the trasnfer destination id
	//
	// required: true
	// example: 2
	Account_destination_id int `json:"account_destination_id"`

	// the transer amount
	//
	// required: true
	// example: 10.20
	Amount float64 `json:"amount"`

	// the account creation date
	//
	// read-only: true
	// example: 2021-05-24T03:30:26.021292Z
	Created_at string `json:"created_at"`
}

//Transfer list
type TransferList struct {
	Transfers []Transfer `json:"transfers"`
}

//Bind method for transfer manipulation
func (t *Transfer) Bind(r *http.Request) error {
	if t.Account_destination_id == 0 || t.Amount == 0.0 {
		return fmt.Errorf("Account destination is a required field")
	}
	return nil
}

//Render method for transfer list
func (*TransferList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

//Render method for transfer
func (*Transfer) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
