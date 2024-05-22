package database

import (
	"database/sql"
	"fmt"
)

// CreateUser создает нового пользователя
// Принимает структуру User со всеми полями, кроме user_id, и создает нового пользователя в бд
// Возвращает идентификатор нового пользователя
func CreateUser(db *sql.DB, u User) (int, error) {
	var userID int

	err := db.QueryRow(`INSERT INTO users(username, email, password, photo, user_role, last_activity, ban, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING user_id`, u.UserName, u.Email, u.Password, u.Photo, u.UserRole, u.LastActivity, u.Ban, u.CreatedAt).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("DBCreateUser: %w", err)
	}

	return userID, nil
}

// GetUserIdByUsername возвращает пользователя по его username
// Принимает username и возвращает структуру User с заполненным полем user_id
func GetUserIdByUsername(db *sql.DB, username string) (User, error) {
	var user User

	row := db.QueryRow(`SELECT user_id FROM users WHERE username = $1`, username)
	err := row.Scan(&user.UserId)
	if err != nil {
		return user, fmt.Errorf("DBGetUserIdByUsername: %w", err)
	}

	return user, nil
}

// GetUserByEmail возвращает пользователя по его email
// Принимает email и возвращает структуру User с данными пользователя
func GetUserByEmail(db *sql.DB, email string) (User, error) {
	var user User

	row := db.QueryRow(`SELECT * FROM users WHERE email = $1`, email)
	err := row.Scan(&user.UserId, &user.UserName, &user.Email, &user.Password, &user.Photo, &user.CreatedAt)
	if err != nil {
		return user, fmt.Errorf("DBGetUserByEmail: %w", err)
	}

	return user, nil
}

// GetUserById возвращает пользователя по его идентификатору
// Принимает user_id и возвращает структуру User с данными пользователя
func GetUserById(db *sql.DB, id int) (User, error) {
	var user User

	row := db.QueryRow(`SELECT * FROM users WHERE user_id = $1`, id)
	err := row.Scan(&user.UserId, &user.UserName, &user.Email, &user.Password, &user.Photo, &user.CreatedAt)
	if err != nil {
		return user, fmt.Errorf("DBGetUserByEmail: %w", err)
	}

	return user, nil
}

// UpdateUserById изменяет данные пользователя
// Принимает структуру User с обновленными данными и идентификатором пользователя
// Возвращает обновленную структуру User
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

// DeleteUserById удаляет пользователя по его идентификатору
// Принимает user_id и удаляет пользователя с этим id
func DeleteUserById(db *sql.DB, id int) error {
	_, err := db.Exec(`DELETE FROM users WHERE user_id = $1`, id)
	if err != nil {
		return fmt.Errorf("DeleteUserById: %w", err)
	}

	return nil
}
