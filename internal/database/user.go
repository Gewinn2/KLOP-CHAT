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

func GetUserIdByUsername(db *sql.DB, username string) (User, error) {
	var user User
	row := db.QueryRow(`SELECT user_id FROM users WHERE username = $1`, username)
	err := row.Scan(&user.Id)
	if err != nil {
		return user, fmt.Errorf("DBGetUserIdByUsername:", err)
	}

	return user, nil
}

func GetUserByEmail(db *sql.DB, email string) (User, error) {
	var user User
	row := db.QueryRow(`SELECT * FROM users WHERE email = $1`, email)
	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.UserName, &user.Photo, &user.CreatedAt)
	if err != nil {
		return user, fmt.Errorf("DBGetUserByEmail:", err)
	}
	return user, nil
}
