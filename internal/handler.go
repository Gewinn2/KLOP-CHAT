package internal

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

func (s *Server) handleConnection(conn net.Conn) { // обработка соединений (они отправляются в слайс buf)
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading from TCP connection:", err)
			break
		}

		request := string(buf[:n])
		fmt.Println("Received request:", request)

		// Парсим запрос
		parts := strings.SplitN(request, " ", 3)
		if len(parts) < 3 {
			fmt.Println("Invalid request format")
			continue
		}
		method := parts[0]
		path := parts[1]
		body := strings.TrimSpace(parts[2])

		// Обработка запроса
		var response string
		switch method {
		case "POST":
			switch path {
			case "/users": // запрос аутентификации (добавления пользователя)
				var user User
				err := json.Unmarshal([]byte(body), &user)
				if err != nil {
					response = "Error parsing user data"
					break
				}

				// TODO: Вот здесь надо сохранять юзера в бд

				response = "User created"

			case "/messages": // запрос отправки сообщения
				var message Message
				err := json.Unmarshal([]byte(body), &message)
				if err != nil {
					response = "Error parsing message data"
					break
				}

				// TODO: тут сохраняем сообщения в бд

				response = "Message created"
			default:
				response = "Invalid path"
			}
		case "GET":
			switch path {
			case "/users/:id": // запрос на получения пользователя по айди

				// TODO: надо вытаскивать пользователя по айди из бд

				response = `{"id": "1", "email": "user@gmal.com", "password": "secret", "user_name": "User", "photo": "photo.jpg", "created_at": "2022-01-01T00:00:00Z"}`

			case "/messages/:id": // запрос на получение сообщения по аййди

				// TODO: вытаскивать пользователя по айди из бд

				response = `{"message_id": "1", "content": "Hello, world!", "user_id": 1, "chat_id": 1, "created_at": "2022-01-01T00:00:00Z"}`

			default:
				response = "Invalid path" // если путь не прокнул
			}
		default:
			response = "Invalid method" // если метод не прокнул (типа гет или пост неправильно написали)
		}

		conn.Write([]byte(response + "\n")) // высылается отчетное сообщение
	}
}
