package userdb

import (
	"fmt"

	"github.com/pkg/errors"
)

func (userDB *UserDB) CreateUser(username string, pw_hash string) error {
	if _, err := userDB.db.Exec(`
		INSERT INTO users (username, pw_hash, created_at, last_login) 
		VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
	`); err != nil {
		return errors.Wrap(err, fmt.Sprintf("Unable to CreateUser with user: %s", username))
	}

	return nil
}
