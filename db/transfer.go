package db

import (
	"github.com/marimell09/stone-challenge/models"
)

func (db Database) GetAllTransfers() (*models.TransferList, error) {
	list := &models.TransferList{}
	rows, err := db.Conn.Query("SELECT * FROM transfers ORDER BY id DESC")

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
