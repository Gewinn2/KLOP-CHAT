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

	_, err = db.Exec(`INSERT INTO chats_participants (chat_id, user_id)
		VALUES ($1, $2)`, chatID, userId1)
	if err != nil {
		return chat, fmt.Errorf("DBCreateChat:", err)
	}

	_, err = db.Exec(`INSERT INTO chats_participants (chat_id, user_id)
		VALUES ($1, $2)`, chatID, userId2)
	if err != nil {
		return chat, fmt.Errorf("DBCreateChat:", err)
	}

	return chat, nil
}

func GetAllChatByUserId(db *sql.DB, userId int) ([]Chat, error) {
	var chatArr []Chat

	var chatIdArr []int
	rows, err := db.Query(`SELECT chat_id FROM chats_participants WHERE user_id = $1`, userId)
	if err != nil {
		return chatArr, fmt.Errorf("DBGetAllChatByUserId:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var chatParticipant ChatParticipant
		err = rows.Scan(&chatParticipant.ChatId)
		if err != nil {
			return chatArr, fmt.Errorf("DBGetAllChatByUserId:", err)
		}
		chatIdArr = append(chatIdArr, chatParticipant.ChatId)
	}

	for i := range chatIdArr {
		var chat Chat
		row := db.QueryRow(`SELECT * FROM chats WHERE chat_id = $1`, chatIdArr[i])
		err = row.Scan(&chat.ChatId, &chat.Name, &chat.Photo, &chat.CreatedAt)
		if err != nil {
			return chatArr, fmt.Errorf("DBGetAllChatByUserId:", err)
		}
		chatArr = append(chatArr, chat)
	}

	return chatArr, nil
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
