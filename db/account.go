package db

import (
	"database/sql"

	"github.com/marimell09/stone-challenge/models"
)

func (db Database) GetAllAccounts() (*models.AccountList, error) {
	list := &models.AccountList{}
	rows, err := db.Conn.Query("SELECT * FROM accounts ORDER BY id DESC")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var account models.Account
		err := rows.Scan(&account.Id, &account.Name, &account.Cpf, &account.Secret, &account.Balance, &account.CreatedAt)
		if err != nil {
			return list, err
		}
		list.Accounts = append(list.Accounts, account)
	}
	return list, nil
}

func (db Database) AddAccount(account *models.Account) error {
	var id int
	var createdAt string
	query := `INSERT INTO accounts (name, cpf, secret, balance) VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	err := db.Conn.QueryRow(query, account.Name, account.Cpf, account.Secret, account.Balance).Scan(&id, &createdAt)
	if err != nil {
		return err
	}

	account.Id = id
	account.CreatedAt = createdAt
	return nil
}

func (db Database) GetAccountById(accountId int) (models.Account, error) {
	account := models.Account{}
	query := `SELECT * FROM accounts WHERE id = $1;`
	row := db.Conn.QueryRow(query, accountId)
	switch err := row.Scan(&account.Id, &account.Name, &account.Cpf, &account.Secret, &account.Balance, &account.CreatedAt); err {
	case sql.ErrNoRows:
		return account, ErrNoMatch
	default:
		return account, err
	}
}

func (db Database) GetAccountBalanceById(accountId int) (float64, error) {
	balance := 0.0
	query := `SELECT balance From accounts WHERE id = $1;`
	row := db.Conn.QueryRow(query, accountId)
	switch err := row.Scan(&balance); err {
	case sql.ErrNoRows:
		return balance, ErrNoMatch
	default:
		return balance, err
	}
}

func (db Database) GetAccountByCpf(accountCpf string) (int, string, error) {
	id := 0
	cpf := ""
	query := `SELECT id, secret FROM accounts WHERE cpf = $1;`
	row := db.Conn.QueryRow(query, accountCpf)
	switch err := row.Scan(&id, &cpf); err {
	case sql.ErrNoRows:
		return id, cpf, ErrNoMatch
	default:
		return id, cpf, err
	}
}

func (db Database) DeleteAccount(accountId int) error {
	query := `DELETE FROM accounts WHERE id = $1;`
	_, err := db.Conn.Exec(query, accountId)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

func (db Database) UpdateAccountBalance(account_id int, balance float64) error {
	id := 0
	query := `UPDATE accounts SET balance=$1 WHERE id=$2 RETURNING id;`
	err := db.Conn.QueryRow(query, balance, account_id).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrNoMatch
		}
		return err
	}
	return nil
}

func (db Database) UpdateAccount(accountId int, accountData models.Account) (models.Account, error) {
	account := models.Account{}
	query := `UPDATE accounts SET name=$1, cpf=$2, secret=$3, balance=$4 WHERE id=$5 RETURNING id, name, cpf, balance, created_at;`
	err := db.Conn.QueryRow(query, accountData.Name, accountData.Cpf, accountData.Secret, accountData.Balance, accountId).Scan(&account.Id, &account.Name, &account.Cpf, &account.Balance, &account.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return account, ErrNoMatch
		}
		return account, err
	}
	return account, nil
}
