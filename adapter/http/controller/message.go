package controller

type Message struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewMessage(status int, message string) *Message {
	return &Message{
		Status:  status,
		Message: message,
	}
}
