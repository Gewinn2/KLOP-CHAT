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
	router.GET("/users/:id", s.getUser)     // получить айди пользователя
	router.POST("/sign-up", s.HandleSignUp) // создание пользователя (запрос для регистрации)
	router.POST("/sign-in", s.HandleSignIn) // создание пользователя (запрос для регистрации)

	authGroup := router.Group("/auth")
	authGroup.Use(pkg.WithJWTAuth)
	{
		authGroup.GET("chat", s.getAllUsersChats)
		authGroup.POST("chat", s.createChat)
		authGroup.PUT("chat", s.updateChat)    // TODO: сделать функцию
		authGroup.DELETE("chat", s.deleteChat) // TODO: сделать функцию

		authGroup.GET("message", s.getAllChatsMessages)
		authGroup.GET("last_message", s.getLatestMessage) // TODO: сделать функцию
		authGroup.POST("message", s.createMessage)
		authGroup.PUT("message", s.updateChat)    // TODO: сделать функцию
		authGroup.DELETE("message", s.deleteChat) // TODO: сделать функцию

	}

	// Запуск сервера
	go func() {
		err := http.Serve(s.listener, router)
		if err != nil {
			fmt.Println("Error starting server:", err)
		}
	}()

	// Прием tcp соединений
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("Error accepting TCP connection:", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) hello(c *gin.Context) {
	c.JSON(http.StatusOK, "Добро пожаловать в чат Клопов!")
}
