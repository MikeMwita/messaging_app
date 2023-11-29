package models

type Message struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Time    string `json:"time"`
	Content string `json:"content"`
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Image string `json:"image"`
}
