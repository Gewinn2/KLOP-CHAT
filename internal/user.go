package internal

import (
	"Astra_Linux_chat/internal/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) createUser(c *gin.Context) { // создаем юзера
	var user database.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.CreateUser(s.DB, user)
	if err != nil {
		fmt.Println("createUser:", err)
	}

	c.JSON(http.StatusCreated, user)
}

func (s *Server) getUser(c *gin.Context) { // получаем юзера по айди
	id := c.Param("id")

	// TODO: надо вытаскивать юзера из бд

	user := database.User{
		Id:        id,
		Email:     "user@example.com",
		Password:  "secret",
		UserName:  "User",
		Photo:     "photo.jpg",
		CreatedAt: "2022-01-01T00:00:00Z",
	}
	c.JSON(http.StatusOK, user)
}
