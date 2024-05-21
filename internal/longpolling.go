package internal

import (
	"Astra_Linux_chat/internal/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"strconv"
)



func (s *Server) longPollMessages(c *gin.Context) {
    // Получаем последнее сообщение, если оно есть
    lastMessageID := c.Query("last_message_id")

    // Ожидаем обновлений сообщений до тех пор, пока не истечет таймаут или не появится новое сообщение
    updates := s.getNewMessages(lastMessageID)
    timeout := time.After(10 * time.Second) // установите таймаут в 10 секунд

    select {
    case <-timeout:
        // Если таймаут истек, вернем пустой ответ
        c.JSON(http.StatusOK, gin.H{})
        return
    case update := <-updates:
        // Если появилось обновление, отправим его клиенту
        c.JSON(http.StatusOK, update)
        return
    }
}


func (s *Server) getNewMessages(lastMessageID string) chan database.Message {
    updates := make(chan database.Message)

    // Запустите горутину, которая будет отправлять обновления в канал
    go func() {
        var latestMessageID int
        if lastMessageID == "" {
            // Если last_message_id не указан, получаем последнее сообщение из базы данных
            var err error
            latestMessageID, err = s.getLatestMessageID()
            if err != nil {
                fmt.Println("Error getting latest message ID:", err)
                return
            }
        } else {
            latestMessageID, _ = strconv.Atoi(lastMessageID)
        }

        for {
            messages, err := database.GetNewMessagesFromDB(s.DB, fmt.Sprintf("%d", latestMessageID))
            if err != nil {
                fmt.Println("Error getting new messages:", err)
                return
            }

            for _, message := range messages {
                updates <- message

                // Обновляем последнее полученное сообщение
                latestMessageID = message.MessageId
            }

            // Ждем некоторое время перед очередным опросом базы данных
            time.Sleep(500 * time.Millisecond)
        }
    }()

    return updates
}

func (s *Server) getLatestMessageID() (int, error) {
    var id int
    err := s.DB.QueryRow("SELECT id FROM messages ORDER BY timestamp DESC LIMIT 1").Scan(&id)
    if err != nil {
        return 0, err
    }
    return id, nil
}
