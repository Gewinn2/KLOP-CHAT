package database

import (
	"Astra_Linux_chat/config"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	db, err := sql.Open("postgres", config.ConnStr)
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
		username VARCHAR(50) NOT NULL,
    	email varchar(100) NOT NULL,
    	password TEXT NOT NULL,
    	photo TEXT NOT NULL,
    	user_role varchar(10) NOT NULL,
    	last_activity TIMESTAMP NOT NULL,
    	ban varchar(20) NOT NULL,
    	created_at TIMESTAMP
	);`)
	if err != nil {
		log.Fatal(err)
	}

	// Создание таблицы чатов
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS chats (
		chat_id SERIAL PRIMARY KEY NOT NULL,
		name VARCHAR(50) NOT NULL,
    	photo TEXT,
    	created_at TIMESTAMP
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
    	created_at TIMESTAMP,
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

	return db
}
