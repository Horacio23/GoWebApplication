package models

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func getDBConnection() (*sql.DB, error) {
	var hostProp string

	if host := os.Getenv("DB"); host != "" {
		hostProp = "host=" + host
		fmt.Println("External host found: ", hostProp)
	} else {
		fmt.Println("No external host found, using default settings")
	}

	db, err := sql.Open("postgres", "dbname=goapp user=postgres password=admin sslmode=disable "+hostProp)
	return db, err
}
