package internal

type User struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UserName  string `json:"user_name"`
	Photo     string `json:"photo"`
	CreatedAt string `json:"created_at"`
}

type Message struct {
	Id        string `json:"message_id"`
	Content   string `json:"content"`
	UserId    int    `json:"user_id"`
	ChatId    int    `json:"chat_id"`
	CreatedAt string `json:"createdat"`
}
