package main

import (
	"Astra_Linux_chat/internal"
	"Astra_Linux_chat/internal/database"
	"fmt"
)

func main() {
	db := database.NewDB()

	server, err := internal.NewServer(8080, db)
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
