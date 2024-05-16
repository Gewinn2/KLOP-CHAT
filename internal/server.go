package internal

import (
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

	// Пользователи
	router.GET("/users/:id", s.getUser)   // получить айди пользователя
	router.POST("/sign-up", s.createUser) // создание пользователя (запрос для аутентификации)

	// Сообщение
	router.GET("/messages/:id", s.getMessage) // получить айди сообщения
	router.POST("/messages", s.createMessage) // псоздать сообщение

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
