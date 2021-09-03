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

// GetUserByUsername gets a user by username, returning an error if not found. `username` should be lowercase string.
func (userDB *UserDB) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	user := &User{}
	if err := userDB.db.QueryRowContext(ctx, getUserByUsernameQuery, username).Scan(&user.Username, &user.DisplayName, &user.Email, &user.LastLogin, &user.CreatedAt); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Unable to GetUserByUsername with username %s", username))
	}

	return user, nil
}

func (userDB *UserDB) UserExists(ctx context.Context, username string) bool {
	username = strings.ToLower(username)
	_, err := userDB.GetUserByUsername(ctx, username)
	return err != nil
}
