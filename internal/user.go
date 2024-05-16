package internal

import (
	"Astra_Linux_chat/internal/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func (s *Server) createUser(c *gin.Context) { // создаем юзера
	var user database.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.CreatedAt = time.Now().Format("2006-01-02")

	_, err := database.UserCheck(s.DB, user.UserName)
	if err == nil {
		c.JSON(http.StatusBadRequest, "Пользователь с таким username уже существует")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("HandleSignUp:", err)
		return
	}

	user.Password = string(hash)

	id, err := database.CreateUser(s.DB, user)
	if err != nil {
		fmt.Println("createUser:", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, id)
}

func (s *Server) getUser(c *gin.Context) { // получаем юзера по айди
	idStr := c.Param("id")
	id,_ := strconv.Atoi(idStr)

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
