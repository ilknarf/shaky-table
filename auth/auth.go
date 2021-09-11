package auth

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

const cookieAuthKeyPath = "./cookie.auth.key"
const cookieEncryptionKeyPath = "./cookie.encryption.key"

type Auth struct {
	store *sessions.CookieStore
}

type User struct {
	username string
	isAdmin  bool
	loggedIn bool
	expires  time.Time
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

func (auth *Auth) withAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s, err := auth.store.Get(r, "login-data")
		if err != nil {
			log.Println("withAuthentication unable to decode session")
			redirectAuthFailed(w, r)
			return
		}

		isLoggedInNoType, ok := s.Values["loggedIn"]
		if !ok {
			log.Println("withAuthentication loggedIn not found on session")
			redirectAuthFailed(w, r)
			return
		}

		isLoggedIn, ok := isLoggedInNoType.(bool)
		if !ok {
			log.Println("withAuthentication loggedIn unable to convert to bool")
			redirectAuthFailed(w, r)
			return
		}

		if !isLoggedIn {
			log.Println("withAuthentication loggedIn false")
			redirectAuthFailed(w, r)
			return
		}

		next(w, r)
	}
}

func redirectAuthFailed(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusUnauthorized)
}
