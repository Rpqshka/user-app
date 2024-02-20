package service

import (
	userApp "user-app"
	"user-app/pkg/repository"
)

type User interface {
	Create(user userApp.User) error
	GetAll() ([]userApp.User, int64, error)
	GetByDateAndAge(input userApp.SearchInput) ([]userApp.User, int64, error)
}

type Service struct {
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
	}
}
