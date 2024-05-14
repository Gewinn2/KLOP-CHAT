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

	// Создание таблицы пользователей
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		user_id SERIAL PRIMARY KEY NOT NULL,
		name VARCHAR(50) NOT NULL,
    	email varchar(100) NOT NULL,
    	password VARCHAR(30) NOT NULL,
    	photo text,
    	created_at VARCHAR(20)
	);`)
	if err != nil {
		log.Fatal(err)
	}

	// Создание таблицы чатов
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS chats (
		chat_id SERIAL PRIMARY KEY NOT NULL,
		name VARCHAR(50) NOT NULL,
    	photo text NOT NULL,
    	created_at VARCHAR(20)
	);`)
	if err != nil {
		log.Fatal(err)
	}

	// Создание таблицы сообщений
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS messages (
		message_id SERIAL PRIMARY KEY NOT NULL,
		content text NOT NULL,
    	chat_id INT NOT NULL,
    	user_id INT NOT NULL,
    	created_at VARCHAR(20),
    	FOREIGN KEY (user_id) REFERENCES users (user_id),
    	FOREIGN KEY (chat_id) REFERENCES chats (chat_id) ON DELETE CASCADE
	);`)
	if err != nil {
		log.Fatal(err)
	}

	// Создание таблицы участников чата
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS chats_partiсipants (
		partiсipants_id SERIAL PRIMARY KEY NOT NULL,
    	chat_id INT NOT NULL,
    	user_id INT NOT NULL,
    	FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE,
    	FOREIGN KEY (chat_id) REFERENCES chats (chat_id) ON DELETE CASCADE
	);`)
	if err != nil {
		log.Fatal(err)
	}

}
