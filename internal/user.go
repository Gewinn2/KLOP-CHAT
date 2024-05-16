package internal

import (
	"Astra_Linux_chat/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (s *Server) getUser(c *gin.Context) { // получаем юзера по айди
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	// TODO: надо вытаскивать юзера из бд

	user := database.User{
		UserId:    id,
		Email:     "user@example.com",
		Password:  "secret",
		UserName:  "User",
		Photo:     "photo.jpg",
		CreatedAt: "2022-01-01T00:00:00Z",
	}
	c.JSON(http.StatusOK, user)
}
