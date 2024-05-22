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
// Функция обновления активности, которая работает по принципу - отправил хттп запрос = онлайн
func (s *Server) userActivity (c *gin.Context){
    userIdToConv := c.Param("user_id")
	userId, _ := strconv.Atoi(userIdToConv)

    // Получаем информацию о пользователе из базы данных
    var user database.User
    if _, err := database.GetUserById(s.DB, userId); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
        return
    }
	// Обновляем время последней активности пользователя в базе данных
    if err := user.UpdateLastActivity(s.DB); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
        return
    }

    // Возвращаем дату последней активности в формате "2022-01-000T00:00:00Z"
    c.JSON(http.StatusOK, gin.H{"last_activity": user.LastActivity})
}
