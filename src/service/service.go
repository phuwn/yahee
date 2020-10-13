package service

import "github.com/phuwn/yahee/src/service/google"

// Service - 3rd parties service handling implementation
type Service struct {
	Google google.Service
}

// New - create new service variable
func New() *Service {
	return &Service{
		Google: google.NewService(),
	}
}
