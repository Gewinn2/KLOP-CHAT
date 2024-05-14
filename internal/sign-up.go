package internal

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UserName  string `json:"user_name"`
	Photo     string `json:"photo"`
	CreatedAt string `json:"created_at"`
}

func HandleSignUp(c *gin.Context) {

}
