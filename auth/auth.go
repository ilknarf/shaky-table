package auth

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/pkg/errors"
)

const (
	cookieAuthKeyPath       = "./cookie.auth.key"
	cookieEncryptionKeyPath = "./cookie.encryption.key"
)

var (
	UsernameNotFoundError = errors.New("username not found on session")
	LoggedInNotFoundError = errors.New("loggedIn not found on session")
	IsAdminNotFoundError  = errors.New("isAdmin not found on session")
)

type Auth struct {
	store *sessions.CookieStore
}

type User struct {
	Username string
	IsAdmin  bool
	LoggedIn bool
}

func readOrGenerateKey(filename string, size int) ([]byte, error) {
	key, err := os.ReadFile(filename)
	if err != nil {
		log.Println("Generating " + filename)
		key = securecookie.GenerateRandomKey(32)
		if err := os.WriteFile(filename, key, 0644); err != nil {
			return nil, err
		}
	}

	return key, nil
}

func New() (*Auth, error) {
	authKey, err := readOrGenerateKey(cookieAuthKeyPath, 64)
	if err != nil {
		return nil, err
	}

	encryptionKey, err := readOrGenerateKey(cookieEncryptionKeyPath, 32)
	if err != nil {
		return nil, err
	}

	return &Auth{sessions.NewCookieStore(authKey, encryptionKey)}, nil
}

// WithAuthentication is middleware that returns an http.HandlerFunc wrapped by auth validatioon
func (auth *Auth) WithAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := auth.GetUser(r)
		if err != nil {
			log.Println(errors.Wrap(err, "withAuthentication unable to authenticate:"))
			redirectAuthFailed(w, r)
			return
		}

		if !u.LoggedIn {
			log.Println(errors.Wrap(err, "withAuthentication user not logged in:"))
			redirectAuthFailed(w, r)
			return
		}

		next(w, r)
	}
}

func (auth *Auth) AddUserToSession(w http.ResponseWriter, r *http.Request, user User) error {
	s, err := auth.store.Get(r, "login-data")
	if err != nil {
		return err
	}

	s.Values["username"] = user.Username
	s.Values["loggedIn"] = user.LoggedIn
	s.Values["isAdmin"] = user.IsAdmin

	s.Save(r, w)

	return nil
}

func (auth *Auth) GetUser(r *http.Request) (*User, error) {
	s, err := auth.store.Get(r, "login-data")
	if err != nil {
		return nil, err
	}

	u := &User{}
	if err != nil {
		return nil, err
	}

	username, ok := s.Values["username"].(string)
	if !ok {
		return nil, UsernameNotFoundError
	}
	u.Username = username

	isAdmin, ok := s.Values["isAdmin"].(bool)
	if !ok {
		return nil, IsAdminNotFoundError
	}

	u.IsAdmin = isAdmin

	loggedIn, ok := s.Values["loggedIn"].(bool)
	if !ok {
		return nil, LoggedInNotFoundError
	}

	u.LoggedIn = loggedIn

	return u, nil
}

func redirectAuthFailed(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusUnauthorized)
}
