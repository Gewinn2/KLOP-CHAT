package database

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
type Chat struct {
	Id        int       `json:"id"`
	UserId1   int       `json:"user_id_1"`
	UserId2   int       `json:"user_id_2"`
	CreatedAt string    `json:"chat_created_at"`
}
