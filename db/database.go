package db

import (
	"database/sql"
	"log"
)

var (
	DB *sql.DB
)

func ResetDB(db *sql.DB) error {
	query := `DROP DATABASE IF EXISTS church`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err.Error())
	}

	return nil
}
