package database

import (
	"Astra_Linux_chat/config"
	"Astra_Linux_chat/pkg"
	"database/sql"
	"fmt"
)

// CreateMessage создает новое сообщение
// Принимает структуру со всеми полями, кроме message_id, и создает новое сообщение в бд
// Возвращает id нового сообщения
func CreateMessage(db *sql.DB, message Message) (int, error) {
	var messageId int

	err := db.QueryRow(`INSERT INTO messages(content, user_id, chat_id, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING message_id`, message.Content, message.UserId, message.ChatId, message.CreatedAt).Scan(&messageId)
	if err != nil {
		return 0, fmt.Errorf("DBCreateMessage: %w", err)
	}
	return messageId, nil
}

// GetLatestMessages возвращает последнее сообщение чата
// Принимает chat_id и возвращает последнее сообщение чата
func GetLatestMessages(db *sql.DB, chatId int) (Message, error) {
	var message Message

	row := db.QueryRow(`SELECT content, created_at FROM messages WHERE chat_id = $1 ORDER BY message_id DESC LIMIT 1`, chatId)
	err := row.Scan(&message.Content, &message.CreatedAt)
	if err != nil {
		return message, fmt.Errorf("GetLatestMessages: %w", err)
	}

	key := []byte(config.EncryptionKey)
	decryptedMessage, err := pkg.DecryptMessage(message.Content, key)
	if err != nil {
		return message, err
	}

	message.Content = decryptedMessage

	return message, nil
}

// GetAllMessages возвращает все сообщения чата
// Принимает chat_id и возвращает массив сообщений этого чата
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

		key := []byte(config.EncryptionKey)
		decryptedMessage, err := pkg.DecryptMessage(message.Content, key)
		if err != nil {
			return messageArr, err
		}

		message.Content = decryptedMessage

		messageArr = append(messageArr, message)
	}

	if err = rows.Err(); err != nil {
		return messageArr, fmt.Errorf("DBGetAllMessages: %w", err)
	}

	return messageArr, nil
}

// UpdateMessageById изменяет сообщение
// Принимает структуру, в которой обязательно должны быть заполнены поля message_id и content
// Изменяет content сообщения, с переданным в структуре id
func UpdateMessageById(db *sql.DB, message Message) (Message, error) {
	_, err := db.Exec(`UPDATE messages 
		SET content = $1
		WHERE message_id = $2`, message.Content, message.MessageId)
	if err != nil {
		return message, fmt.Errorf("DBUpdateMessageById: %w", err)
	}

	return message, nil
}

// DeleteMessageById удаляет сообщение по его id
// Принимает message_id и удаляет сообщение с этим id
func DeleteMessageById(db *sql.DB, messageId int) error {
	_, err := db.Exec(`DELETE FROM messages WHERE message_id = $1`, messageId)
	if err != nil {
		return fmt.Errorf("DBDeleteMessageById: %w", err)
	}

	return nil
}
