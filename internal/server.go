package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

type Server struct {
	listener net.Listener
}

func NewServer(port int) (*Server, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Error connecting to the server")
		return nil, err
	}
	return &Server{listener: listener}, nil
}

func (s *Server) Start() error {
	router := gin.Default()

	router.POST("/users", s.createUser)       // создание пользователя (запрос для аутентификации)
	router.GET("/users/:id", s.getUser)       // получить айди пользователя
	router.POST("/messages", s.createMessage) // псоздать сообщение
	router.GET("/messages/:id", s.getMessage) // получить айди сообщения

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

func (s *Server) createUser(c *gin.Context) { // создаем юзера
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO: надо сохранять юзера в бд

	c.JSON(http.StatusCreated, user)
}

func (s *Server) getUser(c *gin.Context) { // получаем юзера по айди
	id := c.Param("id")

	// TODO: надо вытаскивать юзера из бд

	user := User{
		Id:        id,
		Email:     "user@example.com",
		Password:  "secret",
		UserName:  "User",
		Photo:     "photo.jpg",
		CreatedAt: "2022-01-01T00:00:00Z",
	}
	c.JSON(http.StatusOK, user)
}

func (s *Server) createMessage(c *gin.Context) { // создаем сообщение
	var message Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO: надо сохранять сообщение в бд
	c.JSON(http.StatusCreated, message)
}

func (s *Server) getMessage(c *gin.Context) { // получаем сообщение по айди
	id := c.Param("id")
	// TODO: достаем сообщение из бд
	message := Message{
		Id:        id,
		Content:   "Hi",
		UserId:    5,
		ChatId:    2,
		CreatedAt: "2022-01-01T00:00:00Z",
	}
	c.JSON(http.StatusOK, message)
}


