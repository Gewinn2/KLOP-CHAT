package internal

import (
	"Astra_Linux_chat/internal/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type fullCreateChat struct {
	Name      string `json:"name"`
	Photo     string `json:"photo"`
	UserIdArr []int  `json:"user_id_arr"`
}

type createChatResult struct {
	ChatId         int   `json:"chat_id"`
	ParticipantsId []int `json:"participants_id"`
}

func (s *Server) createChat(c *gin.Context) {
	userIdToConv, ok := c.Get("userId")
	if !ok {
		c.String(http.StatusUnauthorized, "User ID not found")
		fmt.Println("HandleDeleteAdvertisement:", ok)
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

func (s *Server) getAllUsersChats(c *gin.Context) {
	userIdToConv, ok := c.Get("userId")
	if !ok {
		c.String(http.StatusUnauthorized, "User ID not found")
		fmt.Println("HandleDeleteAdvertisement:", ok)
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

func (s *Server) updateChat(c *gin.Context) {

}

func (s *Server) deleteChat(c *gin.Context) {

}
