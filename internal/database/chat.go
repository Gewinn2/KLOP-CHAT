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

type GetAllChatByUserIdResult struct {
	ChatId    int    `json:"chat_id"`
	Name      string `json:"name"`
	Photo     string `json:"photo"`
	Content   string `json:"content"`
	CreatedAt string `json:"message_created_at"`
}

// GetAllChatByUserId возвращает все чаты пользователя
// Принимает user_id и возвращает массив чатов, в которых участвует этот пользователь
func GetAllChatByUserId(db *sql.DB, userId int) ([]GetAllChatByUserIdResult, error) {
	var resultArr []GetAllChatByUserIdResult

	var chatIdArr []int
	rows, err := db.Query(`SELECT chat_id FROM chats_partiсipants WHERE user_id = $1`, userId)
	if err != nil {
		return resultArr, fmt.Errorf("DBGetAllChatByUserId: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var chatParticipant ChatParticipant
		err = rows.Scan(&chatParticipant.ChatId)
		if err != nil {
			return resultArr, fmt.Errorf("DBGetAllChatByUserId: %w", err)
		}
		chatIdArr = append(chatIdArr, chatParticipant.ChatId)
	}

	for i := range chatIdArr {
		var result GetAllChatByUserIdResult

		var chat Chat
		row := db.QueryRow(`SELECT * FROM chats WHERE chat_id = $1`, chatIdArr[i])
		err = row.Scan(&chat.ChatId, &chat.Name, &chat.Photo, &chat.CreatedAt)
		if err != nil {
			return resultArr, fmt.Errorf("DBGetAllChatByUserId: %w", err)
		}
		message, err := GetLatestMessages(db, chatIdArr[i])
		if err != nil {
			return resultArr, fmt.Errorf("DBGetAllChatByUserId: %w", err)
		}
		result.ChatId = chat.ChatId
		result.Name = chat.Name
		result.Photo = chat.Photo
		result.Content = message.Content
		result.CreatedAt = message.CreatedAt
		resultArr = append(resultArr, result)
	}

	return resultArr, nil
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
// Функция для поиска чатов
func FindChatById(db *sql.DB, name string) (int, error) {
	var chatID int
    err := db.QueryRow("SELECT chat_id FROM chats WHERE name = ?", name).Scan(&chatID)
	if err != nil {
		return 0, fmt.Errorf("DBFindChatById: %w", err)
	}
	return chatID, nil
}
