package database

import (
	"database/sql"
	"fmt"
)

func AddChatParticipant(db *sql.DB, chatId, userID int) error {
	_, err := db.Exec(`INSERT INTO chats_partiсipants(chat_id, user_id)
		VALUES ($1, $2)`, chatId, userID)
	if err != nil {
		return fmt.Errorf("DBAddChatParticipant:", err)
	}
	return nil
}

func GetChatParticipantsByChatId(db *sql.DB, chatId int) (ChatParticipant, error) {
	var chatParticipant ChatParticipant

	row := db.QueryRow(`SELECT * FROM chats_partiсipants WHERE chat_id = $1`, chatId)
	err := row.Scan(&chatParticipant.ParticipantId, &chatParticipant.ChatId, &chatParticipant.UserId)
	if err != nil {
		return chatParticipant, fmt.Errorf("DBGetChatsParticipantsByChatId:", err)
	}

	return chatParticipant, nil
}

func DeleteChatParticipant(db *sql.DB, chatId, userID int) error {
	_, err := db.Exec(`DELETE FROM chats_partiсipants WHERE chat_id = $1 AND user_id = $2`, chatId, userID)
	if err != nil {
		return fmt.Errorf("DBDeleteChatParticipant:", err)
	}
	return nil
}
