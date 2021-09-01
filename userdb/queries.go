package userdb

const (
	createUserQuery = `
	INSERT INTO users (username, pw_hash, created_at, last_login) 
	VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
	`
)
