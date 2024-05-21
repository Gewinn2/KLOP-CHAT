package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func GetNewMessagesFromDB(db *sql.DB, lastMessageID string) ([]Message, error) {
	// Получаем все новые сообщения, которые были созданы после последнего полученного сообщения
	rows, err := db.Query("SELECT message_id, content, user_id, chat_id, created_at FROM messages WHERE message_id > ? ORDER BY created_at ASC", lastMessageID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var message Message
		err := rows.Scan(&message.MessageId, &message.Content, &message.UserId, &message.ChatId, &message.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	// Если произошла ошибка, возвращаем ее
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

