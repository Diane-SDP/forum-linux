package models

import (
	"database/sql"
)

var DB *sql.DB

func Open() {
	DB, err := sql.Open("sqlite3", "database.models.DB")
	if err != nil {
		print("erreur ouverture db")
		panic(err)
	}
	defer DB.Close()
}
