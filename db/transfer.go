package db

import (
	"github.com/marimell09/stone-challenge/models"
)

//Get all transfers for the given account id
func (db Database) GetAllTransfers(account_id int) (*models.TransferList, error) {
	list := &models.TransferList{}
	rows, err := db.Conn.Query("SELECT * FROM transfers WHERE account_origin_id = $1 ORDER BY id DESC", account_id)

	if err != nil {
		return list, err
	}

	for rows.Next() {
		var transfer models.Transfer
		err := rows.Scan(&transfer.Id, &transfer.Account_origin_id, &transfer.Account_destination_id, &transfer.Amount, &transfer.Created_at)
		if err != nil {
			return list, err
		}
		list.Transfers = append(list.Transfers, transfer)
	}

	return list, nil
}

//Add new transfer
func (db Database) AddTransfer(transfer *models.Transfer) error {
	var id int
	var created_at string
	query := `INSERT INTO transfers (account_origin_id, account_destination_id, amount) VALUES ($1, $2, $3) RETURNING id, created_at`
	err := db.Conn.QueryRow(query, transfer.Account_origin_id, transfer.Account_destination_id, transfer.Amount).Scan(&id, &created_at)
	if err != nil {
		return err
	}

	transfer.Id = id
	transfer.Created_at = created_at
	return nil
}
