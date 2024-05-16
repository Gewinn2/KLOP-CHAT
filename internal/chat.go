package internal

import (
	"Astra_Linux_chat/internal/database"
)

func (s *Server) createChat(userId1, userId2 int) (int, error) {
	// Создаем новый чат
	chat := database.Chat{
		UserId1:   userId1,
		UserId2:   userId2,
		CreatedAt: "2022-01-01T00:00:00Z",
	}

	// TODO: надо сохранять чат в бд

	return chat.ChatId, nil
}
