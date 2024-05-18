package database

import (
	"database/sql"
	"fmt"
)

func CreateMessage(db *sql.DB, message Message) (Message, error) {
	_, err := db.Exec(`INSERT INTO messages(content, user_id, chat_id, created_at)
		VALUES ($1, $2, $3, $4)`, message.Content, message.UserId, message.ChatId, message.CreatedAt)
	if err != nil {
		return message, fmt.Errorf("DBCreateMessage: %w", err)
	}
	return message, nil
}

func GetAllMessages(db *sql.DB, chatId int) ([]Message, error) {
	var messageArr []Message
	rows, err := db.Query(`SELECT * FROM messages WHERE chat_id = $1`, chatId)
	if err != nil {
		return messageArr, fmt.Errorf("DBGetAllMessages: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var message Message
		if err = rows.Scan(&message.MessageId, &message.Content, &message.UserId, &message.ChatId, &message.CreatedAt); err != nil {
			return messageArr, fmt.Errorf("DBGetAllMessages: %w", err)
		}
		messageArr = append(messageArr, message)
	}

	if err = rows.Err(); err != nil {
		return messageArr, fmt.Errorf("DBGetAllMessages: %w", err)
	}

	return messageArr, nil
}

func UpdateMessageById(db *sql.DB, message Message) (Message, error) {
	_, err := db.Exec(`UPDATE messages 
		SET content = $1
		WHERE message_id = $2`, message.Content, message.MessageId)
	if err != nil {
		return message, fmt.Errorf("DBUpdateMessageById: %w", err)
	}

	return message, nil
}

func DeleteMessageById(db *sql.DB, messageId int) error {
	_, err := db.Exec(`DELETE FROM messages WHERE message_id = $1`, messageId)
	if err != nil {
		return fmt.Errorf("DBDeleteMessageById: %w", err)
	}

	return nil
}
