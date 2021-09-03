package userdb

const (
	getUserTableQuery = `
		SELECT * FROM sqlite_master WHERE name='users' and type='table';
	`
	createUsersTableQuery = `
		CREATE TABLE users (
			id           INTEGER PRIMARY KEY AUTOINCREMENT,
			username     VARCHAR(32) UNIQUE NOT NULL,
			display_name VARCHAR(32) NOT NULL,
			email        VARCHAR(255),
			pw_hash      VARCHAR(64) NOT NULL,
			last_login   INTEGER NOT NULL,
			created_at   INTEGER NOT NULL
		);
		CREATE UNIQUE INDEX ix_users_username ON users (username);
		CREATE INDEX ix_users_display_name ON users (display_name);
		CREATE INDEX ix_users_last_login ON users (last_login);
		CREATE UNIQUE INDEX ix_users_email ON users (email);
	`
	getUserByUsernameQuery = `
		SELECT (username, display_name, email, last_login, created_at) FROM users WHERE username = ?;
	`
	createUserQuery = `
		INSERT INTO users (username, display_name, email, pw_hash, created_at, last_login) 
		VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
	`
)
