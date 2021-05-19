package models

import (
	"fmt"
	"net/http"
)

type Transfer struct {
	Id                     int     `json:"id"`
	Account_origin_id      int     `json:"account_origin_id"`
	Account_destination_id int     `json:"account_destination_id"`
	Amount                 float64 `json:"amount"`
	Created_at             string  `json:"created_at"`
}

type TransferList struct {
	Transfers []Transfer `json:"transfers"`
}

func (t *Transfer) Bind(r *http.Request) error {
	if t.Account_destination_id == 0 || t.Amount == 0.0 {
		return fmt.Errorf("Account destination is a required field")
	}
	return nil
}

func (*TransferList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Transfer) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
