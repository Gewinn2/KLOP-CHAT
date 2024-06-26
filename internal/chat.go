package internal

import (
	"Astra_Linux_chat/internal/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"time"
)

type fullCreateChat struct {
	Name        string   `json:"name"`
	Photo       string   `json:"photo"`
	UsernameArr []string `json:"username_arr"`
}

type createChatResult struct {
	ChatId         int    `json:"chat_id"`
	Name           string `json:"name"`
	Photo          string `json:"photo"`
	ParticipantsId []int  `json:"participants_id"`
}

func (s *Server) createChat(c *gin.Context) {
	userIdToConv, ok := c.Get("userId")
	if !ok {
		c.String(http.StatusUnauthorized, "User ID not found")
		fmt.Println("createChat:", ok)
		return
	}
	userId := userIdToConv.(int)

	var full fullCreateChat
	err := c.BindJSON(&full)
	if err != nil {
		fmt.Println("createChat:", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var UserIdArr []int
	UserIdArr = append(UserIdArr, userId)

	var result createChatResult

	rand.Seed(time.Now().UnixNano())
	colors := []string{"https://images.app.goo.gl/nXB9xCGPQXp1WAaR6", "https://images.app.goo.gl/HCwpsUom368AMP6RA",
		"https://images.app.goo.gl/Rz9McvBXmyLbgZc59", "https://images.app.goo.gl/Q2kwvCjMdYEVMjcX6", "https://images.app.goo.gl/9efoY52cXtYwZzEN9",
		"https://images.app.goo.gl/sjfRcq96ypQ4yhqQ7", "https://images.app.goo.gl/NYUzDxHHfVkovkKS8", "https://images.app.goo.gl/APgjWqDgcEBXpSnT9",
		"https://images.app.goo.gl/mp7eUs5F94trFge5A", "https://images.app.goo.gl/zqwTYw5Lo14FKs7w8", "https://images.app.goo.gl/44pDJPCcdvMDe7oG9"}
	colorChat := colors[rand.Intn(11)]

	var chat database.Chat
	chat.Name = full.Name
	chat.Photo = colorChat
	chat.CreatedAt = time.Now().In(time.FixedZone("UTC+3", 3*60*60)).Format("2006-01-02 15:04")
	chatId, err := database.CreateChat(s.DB, chat)
	if err != nil {
		fmt.Println("createChat:", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	result.Name = full.Name
	result.Photo = colorChat
	result.ChatId = chatId

	for i := range full.UsernameArr {
		user, err := database.GetUserIdByUsername(s.DB, full.UsernameArr[i])
		if err != nil {
			fmt.Println("createChat:", err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		UserIdArr = append(UserIdArr, user.UserId)
	}

	for i := range full.UsernameArr {
		participantId, err := database.AddChatParticipant(s.DB, chatId, UserIdArr[i])
		if err != nil {
			fmt.Println("createChat:", err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		result.ParticipantsId = append(result.ParticipantsId, participantId)
	}

	c.JSON(http.StatusOK, result)
}

// Функция возвращает список чатов, приуроченных к пользователю (работает с помощью id пользователя)
func (s *Server) getAllUsersChats(c *gin.Context) {
	userIdToConv, ok := c.Get("userId")
	if !ok {
		c.String(http.StatusUnauthorized, "User ID not found")
		fmt.Println("getAllUsersChats:", ok)
		return
	}
	userId := userIdToConv.(int)

	AllUsersChats, err := database.GetAllChatByUserId(s.DB, userId)
	if err != nil {
		fmt.Println("getAllUsersChats:", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if AllUsersChats == nil {
		nilArr := [1]int{0}
		c.JSON(http.StatusOK, nilArr)
		return
	}

	c.JSON(http.StatusOK, AllUsersChats)
}

// Функция редактирования чата (фото чата, названия)
func (s *Server) updateChat(c *gin.Context) {
	chatIdToConv, ok := c.Get("chat_id")
	if !ok {
		c.String(http.StatusBadRequest, "Invalid chat ID")
		fmt.Println("updateChat:", ok)
		return
	}
	chatId := chatIdToConv.(int)

	chat := database.Chat{}
	err := c.ShouldBindJSON(&chat)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid chat data")
		fmt.Println("updateChat:", err)
		return
	}
	chat.ChatId = chatId

	_, err = database.UpdateChatById(s.DB, chat)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error updating chat")
		fmt.Println("updateChat:", err)
		return
	}

	c.JSON(http.StatusOK, chat)
}

// Функция удаления чата
func (s *Server) deleteChat(c *gin.Context) {
	chatIdStr := c.Query("id")
	chatId, err := strconv.Atoi(chatIdStr)
	if err != nil {
		fmt.Println("getAllChatsMessages:", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = database.DeleteChatById(s.DB, chatId)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error deleting chat")
		fmt.Println("deleteChat:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Chat deleted"})
}

// Функция для поиска чатов
func (s *Server) findChat(c *gin.Context) {

	name := c.Query("name")
	if name == "" {
		c.String(http.StatusUnauthorized, "Chat ID not found")
		fmt.Println("findChat:", "name is empty")
		return
	}

	chat, err := database.FindChatByName(s.DB, name)
	if err != nil {
		fmt.Println("findChat", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, chat)
}

// Функция для возврата списка чатов по приоритетам
func (s *Server) getChatsPriority(c *gin.Context) {
	userIdToConv, ok := c.Get("userId")
	if !ok {
		c.String(http.StatusUnauthorized, "User ID not found")
		fmt.Println("getChatsPriority:", ok)
		return
	}
	userId := userIdToConv.(int)

	AllUsersChats, err := database.GetAllChatByUserId(s.DB, userId)
	if err != nil {
		fmt.Println("getAllUsersChats:", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	sort.Slice(AllUsersChats, func(i, j int) bool {
		first, _ := database.GetAllMessages(s.DB, AllUsersChats[i].ChatId)

		second, _ := database.GetAllMessages(s.DB, AllUsersChats[j].ChatId)

		return len(first) > len(second)
	})
	if len(AllUsersChats) > 5 {
		c.JSON(http.StatusOK, AllUsersChats[:5])
	} else {
		c.JSON(http.StatusOK, AllUsersChats)
	}
}
