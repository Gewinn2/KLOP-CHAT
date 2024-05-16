package database

import (
	"database/sql"
	"fmt"
)

func CreateUser(db *sql.DB, u User) error {
	_, err := db.Exec(`INSERT INTO users(name, email, password, photo, created_at)
	VALUES ($1, $2, $3, $4)`, u.UserName, u.Email, u.Password, u.Photo, u.CreatedAt)
	if err != nil {
		return fmt.Errorf("CreateUser:", err)
	}

	return nil
}
