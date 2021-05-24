package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	HOST = "database"
	PORT = 5432
)

// ErrNoMatch is returned when we request a row that doesn't exist
var ErrNoMatch = fmt.Errorf("No matching record")

//Database structure for db connection
type Database struct {
	Conn *sql.DB
}

//Initialize database based on the env file information
func Initialize(username, password, database string) (Database, error) {
	db := Database{}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, username, password, database)
	conn, err := sql.Open("postgres", dsn)

	if err != nil {
		return db, err
	}
	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	log.Println("Database connection established")

	return db, nil
}
