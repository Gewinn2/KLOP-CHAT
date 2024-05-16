package database

import (
	"database/sql"
	"fmt"
)

func CreateUser(db *sql.DB, u User) (int, error) {
	var userID int

	err := db.QueryRow(`INSERT INTO users(username, email, password, photo, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING user_id`, u.UserName, u.Email, u.Password, u.Photo, u.CreatedAt).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("DBCreateUser:", err)
	}

	return userID, nil
}

func UserCheck(db *sql.DB, username string) (User, error) {
	var user User
	row := db.QueryRow(`SELECT user_id FROM users WHERE username = $1`, username)
	err := row.Scan(&user.Id)
	if err != nil {
		return user, fmt.Errorf("DBCreateUser:", err)
	}

	return user, nil
}
