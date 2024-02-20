package service

import (
	userApp "user-app"
	"user-app/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user userApp.User) error {
	return s.repo.Create(user)
}

func (s *UserService) GetAll() ([]userApp.User, int64, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetByDateAndAge(input userApp.SearchInput) ([]userApp.User, int64, error) {
	return s.repo.GetByDateAndAge(input)
}
