package model

// User data model
type User struct {
	Base
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
