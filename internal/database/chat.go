package database

import (
	"database/sql"
	"fmt"
)

func CreateChat(db *sql.DB, chat Chat, userId1, userId2 int) (Chat, error) {
	var chatID int

	err := db.QueryRow(`INSERT INTO chats(name, photo, created_at)
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
	_, err := db.Exec(`UPDATE chats 
		SET name = $1, photo = $2 
		WHERE chat_id = $3`, chat.Name, chat.Photo, chat.ChatId)
	if err != nil {
		return chat, fmt.Errorf("DBUpdateChatById:", err)
	}

	return chat, nil
}

func DeleteChatById(db *sql.DB, id int) error {
	_, err := db.Exec(`DELETE FROM chats WHERE chat_id = $1`, id)
	if err != nil {
		return fmt.Errorf("DBDeleteChatById:", err)
	}

	return nil
}
