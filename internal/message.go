package internal

import (
	"Astra_Linux_chat/config"
	"Astra_Linux_chat/internal/database"
	"Astra_Linux_chat/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func (s *Server) createMessage(c *gin.Context) { // создаем сообщение
	userIdToConv, ok := c.Get("userId")
	if !ok {
		c.String(http.StatusUnauthorized, "User ID not found")
		fmt.Println("createMessage:", ok)
		return
	}
	userId := userIdToConv.(int)

	var message database.Message

	err := c.BindJSON(&message)
	if err != nil {
		fmt.Println("createMessage:", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// Шифруем сообщение перед сохранением в базу данных
	key := []byte(config.EncryptionKey)
	encryptedBody, err := pkg.EncryptMessage(message.Content, key)
	if err != nil {
		fmt.Println("createMessage:", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	message.Content = encryptedBody

	message.UserId = userId
	message.CreatedAt = time.Now().In(time.FixedZone("UTC+3", 3*60*60)).Format("2006-01-02 15:04")
	messageId, err := database.CreateMessage(s.DB, message)
	if err != nil {
		fmt.Println("createMessage:", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, messageId)
}

func (s *Server) getAllChatsMessages(c *gin.Context) { // получаем сообщение по айди
	chatIdStr := c.Query("id")
	chatId, err := strconv.Atoi(chatIdStr)
	if err != nil {
		fmt.Println("getAllChatsMessages:", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	AllChatsMessages, err := database.GetAllMessages(s.DB, chatId)
	if err != nil {
		fmt.Println("getAllChatsMessages:", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, AllChatsMessages)
}

func (s *Server) updateMessage(c *gin.Context) {
	var message database.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.String(http.StatusBadRequest, "Invalid request payload")
		fmt.Println("updateMessage:", err)
		return
	}

	// Шифруем сообщение перед сохранением в базу данных
	key := []byte(config.EncryptionKey)
	encryptedBody, err := pkg.EncryptMessage(message.Content, key)
	if err != nil {
		fmt.Println("createMessage:", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	message.Content = encryptedBody

	updatedMessage, err := database.UpdateMessageById(s.DB, message)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error updating message")
		fmt.Println("updateMessage:", err)
		return
	}

	c.JSON(http.StatusOK, updatedMessage)
}

func (s *Server) deleteMessage(c *gin.Context) {
	messageIdStr := c.Query("id")
	messageId, err := strconv.Atoi(messageIdStr)
	if err != nil {
		fmt.Println("getAllChatsMessages:", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = database.DeleteMessageById(s.DB, messageId)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error deleting message")
		fmt.Println("deleteMessage:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message deleted"})
}
