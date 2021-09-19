package userdb

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type User struct {
	Username    string
	DisplayName string
	Email       string
	IsAdmin     bool
	LastLogin   string
	CreatedAt   string
}

func (userDB *UserDB) CreateUser(ctx context.Context, username string, password string, email string) error {
	pwHash, err := hashPassword(password)

	if err != nil {
		return errors.Wrap(err, "Unable to create user due to failed password creation")
	}

	emailString := sql.NullString{
		String: email,
		Valid:  email != "",
	}

	// display_name can be changed, username can't
	displayName := username
	username = strings.ToLower(username)

	if _, err := userDB.db.ExecContext(ctx, createUserQuery, username, displayName, emailString, pwHash); err != nil {
		return errors.Wrap(err, fmt.Sprintf("Unable to create user with user: %s", username))
	}

	return nil
}

// GetUserByUsername gets a user by username, returning an error if not found.
func (userDB *UserDB) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	username = strings.ToLower(username)
	user := &User{}
	var isAdminInt int

	if err := userDB.db.QueryRowContext(ctx, getUserByUsernameQuery, username).Scan(&user.Username, &user.DisplayName, &user.Email, &isAdminInt, &user.LastLogin, &user.CreatedAt); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Unable to GetUserByUsername with username %s", username))
	}

	user.IsAdmin = isAdminInt != 0

	return user, nil
}

func (userDB *UserDB) UserExists(ctx context.Context, username string) (bool, error) {
	username = strings.ToLower(username)
	row := userDB.db.QueryRowContext(ctx, checkUserExistsByUsernameQuery, username)

	var existsBit int

	if err := row.Scan(&existsBit); err != nil {
		return false, err
	}

	return existsBit == 1, nil
}

func (userDB *UserDB) VerifyLogin(ctx context.Context, username string, password string) (*User, error) {
	pwHash, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	u := &User{}
	var isAdminInt int

	if err := userDB.db.QueryRowContext(ctx, getUserAndVerifyLogin, username, pwHash).Scan(&u.Username, &u.DisplayName, &u.Email, &isAdminInt, &u.LastLogin, &u.CreatedAt); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Unable to VerifyLogin with username %s", username))
	}

	u.IsAdmin = isAdminInt != 0

	return u, nil
}
