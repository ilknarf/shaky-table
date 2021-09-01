package userdb

import (
	"fmt"

	"github.com/pkg/errors"
)

func (userDB *UserDB) CreateUser(username string, password string) error {
	pwHash, err := hashPassword(password)
	if err != nil {
		return errors.Wrap(err, "Unable to CreateUser due to failed password hashing")
	}

	if _, err := userDB.db.Exec(createUserQuery, username, pwHash); err != nil {
		return errors.Wrap(err, fmt.Sprintf("Unable to CreateUser with user: %s", username))
	}

	return nil
}
