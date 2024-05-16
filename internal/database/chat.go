package database

import (
	"database/sql"
	"fmt"
)

func CreateChat(db *sql.DB, chat Chat, userId1, userId2 int) (Chat, error) {
	var chatID int

	err := db.QueryRow(`INSERT INTO chat(name, photo, created_at)
		VALUES ($1, $2, $3)
		RETURNING chat_id`, chat.Name, chat.Photo, chat.CreatedAt).Scan(&chatID)
	if err != nil {
		return chat, fmt.Errorf("DBCreateChat:", err)
	}

	_, err = db.Exec(`INSERT INTO chats_partiсipants (chat_id, user_id)
		VALUES ($1, $2)`, chatID, userId1)
	if err != nil {
		return chat, fmt.Errorf("DBCreateChat:", err)
	}

	_, err = db.Exec(`INSERT INTO chats_partiсipants (chat_id, user_id)
		VALUES ($1, $2)`, chatID, userId2)
	if err != nil {
		return chat, fmt.Errorf("DBCreateChat:", err)
	}

	return chat, nil
}

func UpdateChatById(db *sql.DB, chat Chat) (Chat, error) {
	return chat, nil
}

func DeleteChatById(db *sql.DB) error {
	return nil
}
