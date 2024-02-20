package repository

import (
	"gorm.io/gorm"
	userApp "user-app"
)

type User interface {
	Create(user userApp.User) error
	GetAll() ([]userApp.User, int64, error)
	GetByDateAndAge(input userApp.SearchInput) ([]userApp.User, int64, error)
}

type Repository struct {
	User
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
	}
}
