package userdb

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

type UserDB struct {
	db *sql.DB
}

func createUsersTable(db *sql.DB) error {
	_, err := db.Exec(createUsersTableQuery)

	return err
}

func Open() (*UserDB, error) {
	db, err := sql.Open("sqlite3", "./db/users.db")
	if err != nil {
		return nil, errors.Wrap(err, "Unable to connect to UserDB")
	}

	rows, err := db.Query("SELECT * FROM sqlite_master WHERE name='users' and type='table';")
	if err != nil {
		return nil, errors.Wrap(err, "Unable to read sqlite_master from UserDB")
	}

	if !rows.Next() {
		log.Println("No users table found. Creating new users table.")
		if err := createUsersTable(db); err != nil {
			return nil, errors.Wrap(err, "Unable to create users table")
		}
	}

	return &UserDB{db}, nil
}

func (userDB *UserDB) Close() {
	userDB.db.Close()
}
