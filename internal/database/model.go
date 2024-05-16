package database

type User struct {
	UserId    int    `json:"user_id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UserName  string `json:"username"`
	Photo     string `json:"photo"`
	CreatedAt string `json:"created_at"`
}

type Message struct {
	MessageId int    `json:"message_id"`
	Content   string `json:"content"`
	UserId    int    `json:"user_id"`
	ChatId    int    `json:"chat_id"`
	CreatedAt string `json:"created_at"`
}

type Chat struct {
	ChatId    int    `json:"chat_id"`
	Name      string `json:"name"`
	Photo     string `json:"photo"`
	CreatedAt string `json:"chat_created_at"`
}

type ChatParticipant struct {
	ParticipantId int `json:"participant_id"`
	ChatId        int `json:"chat_id"`
	UserId        int `json:"user_id"`
}
