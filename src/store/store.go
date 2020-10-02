package store

import "github.com/phuwn/yahee/src/store/user"

// Store - server store struct
type Store struct {
	User user.Store
}

// New - create new store variable
func New() *Store {
	return &Store{
		User: user.NewStore(),
	}
}
