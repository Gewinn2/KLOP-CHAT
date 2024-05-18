package database

import (
	"database/sql"
	"fmt"
)

// CreateChat создает новый чат
// Принимает структуру Chat
// Возвращает id нового чата
func CreateChat(db *sql.DB, chat Chat) (int, error) {
	var chatId int

	err := db.QueryRow(`INSERT INTO chats(name, photo, created_at)
		VALUES ($1, $2, $3)
		RETURNING chat_id`, chat.Name, chat.Photo, chat.CreatedAt).Scan(&chatId)
	if err != nil {
		return 0, fmt.Errorf("DBCreateChat: %w", err)
	}

	return chatId, nil
}

// GetAllChatByUserId возвращает все чаты пользователя
// Принимает user_id и возвращает массив чатов, в которых участвует этот пользователь
func GetAllChatByUserId(db *sql.DB, userId int) ([]Chat, error) {
	var chatArr []Chat

	var chatIdArr []int
	rows, err := db.Query(`SELECT chat_id FROM chats_participants WHERE user_id = $1`, userId)
	if err != nil {
		return chatArr, fmt.Errorf("DBGetAllChatByUserId: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var chatParticipant ChatParticipant
		err = rows.Scan(&chatParticipant.ChatId)
		if err != nil {
			return chatArr, fmt.Errorf("DBGetAllChatByUserId: %w", err)
		}
		chatIdArr = append(chatIdArr, chatParticipant.ChatId)
	}

	for i := range chatIdArr {
		var chat Chat
		row := db.QueryRow(`SELECT * FROM chats WHERE chat_id = $1`, chatIdArr[i])
		err = row.Scan(&chat.ChatId, &chat.Name, &chat.Photo, &chat.CreatedAt)
		if err != nil {
			return chatArr, fmt.Errorf("DBGetAllChatByUserId: %w", err)
		}
		chatArr = append(chatArr, chat)
	}

	return chatArr, nil
}

// UpdateChatById изменяет данные чата
// Принимает структуру Chat с обновленными данными чата и идентификатором чата
// Возвращает обновленную структуру Chat
func UpdateChatById(db *sql.DB, chat Chat) (Chat, error) {
	_, err := db.Exec(`UPDATE chats 
		SET name = $1, photo = $2 
		WHERE chat_id = $3`, chat.Name, chat.Photo, chat.ChatId)
	if err != nil {
		return chat, fmt.Errorf("DBUpdateChatById: %w", err)
	}

	return chat, nil
}

// DeleteChatById удаляет чат по его идентификатору
// Принимает chat_id и удаляет чат с этим id
func DeleteChatById(db *sql.DB, id int) error {
	_, err := db.Exec(`DELETE FROM chats WHERE chat_id = $1`, id)
	if err != nil {
		return fmt.Errorf("DBDeleteChatById: %w", err)
	}

	return nil
}
