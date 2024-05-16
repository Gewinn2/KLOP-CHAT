package database

import "database/sql"

func CreateChat(db *sql.DB, chat Chat) (Chat, error) {
	return chat, nil
}

func UpdateChatById(db *sql.DB, chat Chat) (Chat, error) {
	return chat, nil
}

func DeleteChatById(db *sql.DB) error {
	return nil
}
