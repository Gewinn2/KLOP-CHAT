package internal

import (
	"Astra_Linux_chat/internal/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type fullCreateChat struct {
	Name      string `json:"name"`
	Photo     string `json:"photo"`
	UserIdArr []int  `json:"user_id_arr"`
}

type createChatResult struct {
	ChatId         int    `json:"chat_id"`
	Name           string `json:"name"`
	Photo          string `json:"photo"`
	ParticipantsId []int  `json:"participants_id"`
}

func (s *Server) createChat(c *gin.Context) {
	userIdToConv, ok := c.Get("userId")
	if !ok {
		c.String(http.StatusUnauthorized, "User ID not found")
		fmt.Println("createChat:", ok)
		return
	}
	userId := userIdToConv.(int)

	var full fullCreateChat
	err := c.BindJSON(&full)
	if err != nil {
		fmt.Println("createChat:", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	full.UserIdArr = append(full.UserIdArr, userId)

	var result createChatResult

	var chat database.Chat
	chat.Name = full.Name
	chat.Photo = full.Photo
	chat.CreatedAt = time.Now().Format("2006-01-02")
	chatId, err := database.CreateChat(s.DB, chat)
	if err != nil {
		fmt.Println("createChat:", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	result.Name = full.Name
	result.Photo = full.Photo
	result.ChatId = chatId

	for i := range full.UserIdArr {
		participantId, err := database.AddChatParticipant(s.DB, chatId, full.UserIdArr[i])
		if err != nil {
			fmt.Println("createChat:", err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		result.ParticipantsId = append(result.ParticipantsId, participantId)
	}

	c.JSON(http.StatusOK, result)
}

// Функция возвращает список чатов, приуроченных к пользователю (работает с помощью id пользователя)
func (s *Server) getAllUsersChats(c *gin.Context) {
	userIdToConv, ok := c.Get("userId")
	if !ok {
		c.String(http.StatusUnauthorized, "User ID not found")
		fmt.Println("getAllUsersChats:", ok)
		return
	}
	userId := userIdToConv.(int)

	AllUsersChats, err := database.GetAllChatByUserId(s.DB, userId)
	if err != nil {
		fmt.Println("getAllUsersChats:", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, AllUsersChats)
}

// Функция редактирования чата (фото чата, названия)
func (s *Server) updateChat(c *gin.Context) {
	chatIdToConv, ok := c.Get("chat_id")
	if !ok {
		c.String(http.StatusBadRequest, "Invalid chat ID")
		fmt.Println("updateChat:", ok)
		return
	}
	chatId := chatIdToConv.(int)

	chat := database.Chat{}
	err := c.ShouldBindJSON(&chat)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid chat data")
		fmt.Println("updateChat:", err)
		return
	}
	chat.ChatId = chatId

	_, err = database.UpdateChatById(s.DB, chat)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error updating chat")
		fmt.Println("updateChat:", err)
		return
	}

	c.JSON(http.StatusOK, chat)
}

// Функция удаления чата
func (s *Server) deleteChat(c *gin.Context) {
	chatIdStr := c.Query("id")
	chatId, err := strconv.Atoi(chatIdStr)
	if err != nil {
		fmt.Println("getAllChatsMessages:", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = database.DeleteChatById(s.DB, chatId)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error deleting chat")
		fmt.Println("deleteChat:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Chat deleted"})
}

// Функция для поиска чатов
func (s *Server) findChat(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.String(http.StatusUnauthorized, "Chat ID not found")
		fmt.Println("findChat:", "name is empty")
		return
	}

	ChatId, err := database.FindChatByName(s.DB, name)
	if err != nil {
		fmt.Println("findChat", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ChatId)
}

// Функция для возврата списка чатов по приоритетам
func (s *Server) getChatsPriority(c *gin.Context) {
	userIdToConv, ok := c.Get("userId")
	if !ok {
		c.String(http.StatusUnauthorized, "User ID not found")
		fmt.Println("getChatsPriority:", ok)
		return
	}
	userId := userIdToConv.(int)

	AllUsersChats, err := database.GetAllChatByUserId(s.DB, userId)
	if err != nil {
		fmt.Println("getAllUsersChats:", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	sort.Slice(AllUsersChats, func(i, j int) bool {
		first, _ := database.GetAllMessages(s.DB, AllUsersChats[i].ChatId)

		second, _ := database.GetAllMessages(s.DB, AllUsersChats[j].ChatId)

		return len(first) > len(second)
	})
	if len(AllUsersChats) > 5 {
		c.JSON(http.StatusOK, AllUsersChats[:5])
	} else {
		c.JSON(http.StatusOK, AllUsersChats)
	}
}
