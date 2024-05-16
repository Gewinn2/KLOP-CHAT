package internal

import (
	"Astra_Linux_chat/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (s *Server) createMessage(c *gin.Context) { // создаем сообщение
	var message database.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO: надо сохранять сообщение в бд
	c.JSON(http.StatusCreated, message)
}

func (s *Server) getMessage(c *gin.Context) { // получаем сообщение по айди
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	// TODO: достаем сообщение из бд
	message := database.Message{
		Id:        id,
		Content:   "Hi",
		UserId:    5,
		ChatId:    2,
		CreatedAt: "2022-01-01T00:00:00Z",
	}
	c.JSON(http.StatusOK, message)
}
