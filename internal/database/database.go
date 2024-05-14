package database

import (
	AstraLinux_TCPChat_hackathon "Astra_Linux_chat"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewDB() {
	db, err := sql.Open("postgres", AstraLinux_TCPChat_hackathon.ConnStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
