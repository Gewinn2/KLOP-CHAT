package internal

import (
	"Astra_Linux_chat/internal/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Функция получения всех пользователей чата по айди чата
func (s *Server) getChatParticipants(c *gin.Context) {
	chatIdToConv, ok := c.Get("chat_id")
	if !ok {
		c.String(http.StatusBadRequest, "Invalid chat ID")
		fmt.Println("getChatParticipants:", ok)
		return
	}
	chatId := chatIdToConv.(int)

	participants, err := database.GetChatParticipantsByChatId(s.DB, chatId)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error getting chat participants")
		fmt.Println("getChatParticipants:", err)
		return
	}

	c.JSON(http.StatusOK, participants)
}

// Функция добавления пользователя в чат по айди чата и айди пользователя
func (s *Server) addParticipant(c *gin.Context) {
	chatIdToConv, ok := c.Get("chat_id")
	if !ok {
		c.String(http.StatusBadRequest, "Invalid chat ID")
		fmt.Println("addChatParticipant:", ok)
		return
	}
	chatId := chatIdToConv.(int)

	userIdToConv, ok := c.Get("user_id")
	if !ok {
		c.String(http.StatusBadRequest, "Invalid user ID")
		fmt.Println("addChatParticipant:", ok)
		return
	}
	userId := userIdToConv.(int)

	participantsId, err := database.AddChatParticipant(s.DB, chatId, userId)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error adding chat participant")
		fmt.Println("addChatParticipant:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"participants_id": participantsId})
}

// Функция удаления пользователя из чата по айди пользователя
func (s *Server) deleteParticipant(c *gin.Context) {
	participantsIdToConv, ok := c.Get("participants_id")
	if !ok {
		c.String(http.StatusBadRequest, "Invalid participants ID")
		fmt.Println("deleteChatParticipant:", ok)
		return
	}
	participantsId := participantsIdToConv.(int)

	err := database.DeleteChatParticipant(s.DB, participantsId)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error deleting chat participant")
		fmt.Println("deleteChatParticipant:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Chat participant deleted"})
}
