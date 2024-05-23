package internal

import (
	"Astra_Linux_chat/internal/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func (s *Server) HandleSignUp(c *gin.Context) {
	var user database.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.CreatedAt = time.Now().In(time.FixedZone("UTC+3", 3*60*60)).Format("2006-01-02 15:04")
	user.LastActivity = time.Now().In(time.FixedZone("UTC+3", 3*60*60)).Format("2006-01-02 15:04")
	user.Ban = "false"
	user.UserRole = "user"

	_, err := database.GetUserIdByUsername(s.DB, user.UserName)
	if err == nil {
		c.JSON(http.StatusBadRequest, "A user with this username already exists")
		return
	}

	_, err = database.GetUserByEmail(s.DB, user.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, "A user with this email already exists")
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
