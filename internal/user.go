package internal

import (
	"Astra_Linux_chat/internal/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (s *Server) getUserById(c *gin.Context) { // получаем юзера по айди
	userIdStr := c.Query("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		fmt.Println("getUserById:", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	user, err := database.GetUserById(s.DB, userId)
	if err != nil {
		fmt.Println("getUserById:", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
