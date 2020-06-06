package database

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB {

	db, err := sql.Open("mysql", "root:@127.0.0.1/learngolang")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
