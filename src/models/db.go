package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func getDBConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "dbname=mydb user=postgres password=horacio01 sslmode=disable")
	return db, err
}
