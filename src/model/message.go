package model

// Message data model
type Message struct {
	Base
	ID         string `json:"id"`
	SenderID   string `json:"-"`
	ReceiverID string `json:"-"`
	Content    string `json:"content"`
}
