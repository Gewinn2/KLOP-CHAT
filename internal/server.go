package internal

import (
	"Astra_Linux_chat/pkg"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

type Server struct {
	listener net.Listener
	DB       *sql.DB
}

func NewServer(port int, db *sql.DB) (*Server, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Error connecting to the server")
		return nil, err
	}
	return &Server{listener: listener, DB: db}, nil
}

func (s *Server) Start() error {
	router := gin.Default()

	router.Use(pkg.CORSMiddleware())

	router.GET("/", s.hello)

	// Пользователи
	router.GET("/user", s.getUserById)      // получить пользователя по айди
	router.POST("/sign-up", s.HandleSignUp) // создание пользователя (запрос для регистрации)
	router.POST("/sign-in", s.HandleSignIn) // создание пользователя (запрос для регистрации)

	authGroup := router.Group("/auth")
	authGroup.Use(pkg.WithJWTAuth)
	{
		authGroup.GET("participant", s.getChatParticipants)
		authGroup.POST("participant", s.addParticipant)
		authGroup.DELETE("participant", s.deleteParticipant)

		authGroup.GET("chat", s.getAllUsersChats)
		authGroup.GET("chat/priority", s.getChatsPriority)
		authGroup.POST("chat", s.createChat)
		authGroup.PUT("chat", s.updateChat)
		authGroup.DELETE("chat", s.deleteChat)
		authGroup.GET("chat/find", s.findChat)

		authGroup.GET("message", s.getAllChatsMessages)
		authGroup.POST("message", s.createMessage)
		authGroup.PUT("message", s.updateMessage)
		authGroup.DELETE("message", s.deleteMessage)
		authGroup.GET("/message/longpoll", s.longPollMessages)
	}

	// Запуск сервера
	if err := http.Serve(s.listener, router); err != nil {
		fmt.Println("Error starting server:", err)
		return err
	}
	return nil
}

func (s *Server) hello(c *gin.Context) {
	c.JSON(http.StatusOK, "Добро пожаловать в чат Клопов!")
}
