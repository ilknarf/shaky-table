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

func Open() (*UserDB, error) {
	db, err := sql.Open("sqlite3", "./db/users.db")
	if err != nil {
		return nil, errors.Wrap(err, "Unable to connect to UserDB")
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "Unable to begin createDB transaction on UsersDB")
	}

	defer tx.Rollback()

	rows, err := tx.Query(getUserTableQuery)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to read sqlite_master from UserDB")
	}

	if !rows.Next() {
		log.Println("No users table found. Creating new users table.")
		if _, err := tx.Exec(createUsersTableQuery); err != nil {
			return nil, errors.Wrap(err, "Unable to create users table")
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, errors.Wrap(err, "Unable to commit create table txn")
	}

	return &UserDB{db}, nil
}

func (userDB *UserDB) Close() {
	userDB.db.Close()
}
