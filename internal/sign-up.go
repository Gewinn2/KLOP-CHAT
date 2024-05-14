package internal

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	UID      string `json:"uid"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

func HandleSignUp(c *gin.Context) {
	client, err := InitFirebase().Auth(context.Background())
	if err != nil {
		fmt.Println("HandleSignUp:", err)
	}

	var user User

	err = c.BindJSON(&user)
	if err != nil {
		fmt.Println("HandleSignUp:", err)
	}

	params := (&auth.UserToCreate{}).
		Email(user.Email).
		EmailVerified(false).
		Password(user.Password).
		DisplayName(user.Name + " " + user.Surname)

	u, err := client.CreateUser(context.Background(), params)
	if err != nil {
		fmt.Println("HandleSignUp:", err)
		c.JSON(http.StatusForbidden, "Пользователь с таким email уже существует")
	}

	c.JSON(http.StatusCreated, u)
}
