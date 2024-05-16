package internal

import (
	"Astra_Linux_chat/internal/database"
)

func (s *Server) createChat(userId1, userId2 int) (int, error) {
	// Создаем новый чат
	chat := database.Chat{
		ChatId:    0,
		Name:      "",
		Photo:     "",
		CreatedAt: "2022-01-01T00:00:00Z",
	}

	// TODO: надо сохранять чат в бд

	return chat.ChatId, nil
}
