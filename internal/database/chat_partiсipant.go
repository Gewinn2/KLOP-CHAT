package database

import (
	"database/sql"
	"fmt"
)

func AddChatParticipant(db *sql.DB, chatId, userID int) error {
	_, err := db.Exec(`INSERT INTO chats_participants(chat_id, user_id)
		VALUES ($1, $2)`, chatId, userID)
	if err != nil {
		return fmt.Errorf("DBAddChatParticipant: %w", err)
	}
	return nil
}

func GetChatParticipantsByChatId(db *sql.DB, chatId int) ([]ChatParticipant, error) {
	var participants []ChatParticipant

	rows, err := db.Query(`SELECT * FROM chats_participants WHERE chat_id = $1`, chatId)
	if err != nil {
		return participants, fmt.Errorf("DBGetChatParticipantsByChatId: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var participant ChatParticipant
		if err := rows.Scan(&participant.ParticipantId, &participant.ChatId, &participant.UserId); err != nil {
			return participants, fmt.Errorf("DBGetChatParticipantsByChatId: %w", err)
		}
		participants = append(participants, participant)
	}

	if err = rows.Err(); err != nil {
		return participants, fmt.Errorf("DBGetChatParticipantsByChatId: %w", err)
	}

	return participants, nil
}

func DeleteChatParticipant(db *sql.DB, chatId, userID int) error {
	_, err := db.Exec(`DELETE FROM chats_participants WHERE chat_id = $1 AND user_id = $2`, chatId, userID)
	if err != nil {
		return fmt.Errorf("DBDeleteChatParticipant: %w", err)
	}
	return nil
}
