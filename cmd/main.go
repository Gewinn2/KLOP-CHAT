package main

import (
	"Astra_Linux_chat/internal"
	"Astra_Linux_chat/internal/database"
	"Astra_Linux_chat/pkg"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.NewDB()

	r := gin.Default()

	r.Use(pkg.CORSMiddleware())

	r.POST("/sign-up", internal.HandleSignUp)
	r.GET("/", internal.HandleSignIn)

	err := r.Run(":80")
	if err != nil {
		log.Fatal(err)
	}
	server, err := NewServer(8080)
	if err != nil {
		fmt.Println("Error creating server:", err)
		return
	}
	err = server.Start()
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
