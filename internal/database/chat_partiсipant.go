package database

import (
	"database/sql"
	"fmt"
)

// AddChatParticipant добавляет участника в чат
// Принимает chat_id и user_id, создает новую запись в таблице chats_participants и возвращает participants_id новой записи
func AddChatParticipant(db *sql.DB, chatId, userID int) (int, error) {
	var participantsId int
	err := db.QueryRow(`INSERT INTO chats_partiсipants(chat_id, user_id)
		VALUES ($1, $2)
		RETURNING partiсipants_id`, chatId, userID).Scan(&participantsId)
	if err != nil {
		return 0, fmt.Errorf("DBAddChatParticipant: %w", err)
	}
	return participantsId, nil
}

// GetChatParticipantsByChatId возвращает всех участников чата
// Принимает chat_id и возвращает массив участников этого чата
func GetChatParticipantsByChatId(db *sql.DB, chatId int) ([]ChatParticipant, error) {
	var participants []ChatParticipant

	rows, err := db.Query(`SELECT * FROM chats_partiсipants WHERE chat_id = $1`, chatId)
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

// DeleteChatParticipant удаляет участника чата по его id
// Принимает participants_id и удаляет запись с этим id из таблицы chats_participants
func DeleteChatParticipant(db *sql.DB, participantsId int) error {
	_, err := db.Exec(`DELETE FROM chats_partiсipants WHERE partiсipants_id = $1`, participantsId)
	if err != nil {
		return fmt.Errorf("DBDeleteChatParticipant: %w", err)
	}
	return nil
}
