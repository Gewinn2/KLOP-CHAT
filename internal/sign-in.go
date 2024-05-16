package internal

import (
	"Astra_Linux_chat/internal/database"
	"Astra_Linux_chat/pkg"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) HandleSignIn(c *gin.Context) {
	var user database.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// искать юзера в по имейлу (отсылвает на функцию бд)
	foundUser, err := database.GetUserByEmail(s.DB, user.Email)
	if err != nil {
		fmt.Println("HandleSignIn: ", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// сравнить пароли на соответствие
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		fmt.Println("HandleSignIn: ", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// генерируем jwt токен
	token, err := pkg.GenerateJWT(foundUser.Id)
	if err != nil {
		fmt.Println("HandleSignIn: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// присваиваем Jwt токен
	c.Header("Authorization", fmt.Sprintf("Bearer %s", token))

	c.JSON(http.StatusOK, foundUser)
}
