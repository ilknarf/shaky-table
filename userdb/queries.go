package userdb

const (
	createUsersTableQuery = `
		CREATE TABLE users (
			id           INTEGER PRIMARY KEY AUTOINCREMENT,
			username     VARCHAR(32) UNIQUE NOT NULL,
			display_name VARCHAR(32) UNIQUE NOT NULL,
			email        VARCHAR(255),
			pw_hash      VARCHAR(64) NOT NULL,
			last_login   INTEGER NOT NULL,
			created_at   INTEGER NOT NULL
		);
	`
	getUserByUsernameQuery = `
		SELECT (username, display_name, email, last_login, created_at) FROM users WHERE username = ?;
	`
	createUserQuery = `
		INSERT INTO users (username, display_name, email, pw_hash, created_at, last_login) 
		VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
	`
)
