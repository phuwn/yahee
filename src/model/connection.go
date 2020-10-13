package model

type ConnectionStatus int

const (
	NewConn ConnectionStatus = iota
	ConnectedConn
	BLockedConn
)

// Connection data model
type Connection struct {
	Base
	UserID1 string           `json:"-"`
	UserID2 string           `json:"-"`
	Status  ConnectionStatus `json:"status"`
}
