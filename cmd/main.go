package main

import (
	"Astra_Linux_chat/internal"
	"Astra_Linux_chat/pkg"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	r.Use(pkg.CORSMiddleware())

	r.POST("/sign-up", internal.HandleSignUp)
	r.POST("/sign-in", internal.HandleSignIn)

	err := r.Run(":80")
	if err != nil {
		log.Fatal(err)
	}
}
