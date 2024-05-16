package database

import "database/sql"

func CreateMessage(db *sql.DB, message Message) (Message, error) {
	return message, nil
}

func UpdateMessageById(db *sql.DB, message Message) (Message, error) {
	return message, nil
}

func DeleteMessageById(db *sql.DB) error {
	return nil
}
