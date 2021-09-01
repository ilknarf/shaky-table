package userdb

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "db/users.db")
	if err != nil {
		return nil, errors.Wrap("Unable to connect to UserDB", err)
	}

	rows, err := db.Query("SELECT * FROM sqlite_master WHERE name='users' and type='table")
	if err != nil {
		return nil, errors.Wrap("Unable to read sqlite_master from UserDB", err)
	}

	if !rows.Next() {
		log.Println("No users table found. Creating new users table.")
	}

	return db, nil
}
