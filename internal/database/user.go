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
		return 0, fmt.Errorf("DBCreateUser: %w", err)
	}

	return userID, nil
}

func GetUserIdByUsername(db *sql.DB, username string) (User, error) {
	var user User

	row := db.QueryRow(`SELECT user_id FROM users WHERE username = $1`, username)
	err := row.Scan(&user.UserId)
	if err != nil {
		return user, fmt.Errorf("DBGetUserIdByUsername: %w", err)
	}

	return user, nil
}

func GetUserByEmail(db *sql.DB, email string) (User, error) {
	var user User

	row := db.QueryRow(`SELECT * FROM users WHERE email = $1`, email)
	err := row.Scan(&user.UserId, &user.UserName, &user.Email, &user.Password, &user.Photo, &user.CreatedAt)
	if err != nil {
		return user, fmt.Errorf("DBGetUserByEmail: %w", err)
	}

	return user, nil
}

func GetUserById(db *sql.DB, id int) (User, error) {
	var user User

	row := db.QueryRow(`SELECT * FROM users WHERE user_id = $1`, id)
	err := row.Scan(&user.UserId, &user.UserName, &user.Email, &user.Password, &user.Photo, &user.CreatedAt)
	if err != nil {
		return user, fmt.Errorf("DBGetUserByEmail: %w", err)
	}

	return user, nil
}

func UpdateUserById(db *sql.DB, user User) (User, error) {
	_, err := db.Exec(`UPDATE users 
		SET username = $1, 
		    email = $2,
		    password = $3,
		    photo = $4
		WHERE user_id = $5`, user.UserName, user.Email, user.Password, user.Photo, user.UserId)

	if err != nil {
		return user, fmt.Errorf("UpdateUserById: %w", err)
	}

	return user, nil
}

func DeleteUserById(db *sql.DB, id int) error {
	_, err := db.Exec(`DELETE FROM users WHERE user_id = $1`, id)
	if err != nil {
		return fmt.Errorf("DeleteUserById: %w", err)
	}

	return nil
}
