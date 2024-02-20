package repository

import (
	"errors"
	"gorm.io/gorm"
	userApp "user-app"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) Create(user userApp.User) error {
	if err := r.db.First(&user, "id = ?", user.Id).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	} else {
		return errors.New("this user already created")
	}

	if err := r.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserPostgres) GetAll() ([]userApp.User, int64, error) {
	var users []userApp.User

	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	count := result.RowsAffected
	return users, count, nil

}

func (r *UserPostgres) GetByDateAndAge(input userApp.SearchInput) ([]userApp.User, int64, error) {
	var users []userApp.User

	query := r.db.Where("recording_date BETWEEN ? AND ?", input.StartDate, input.EndDate).
		Where("age BETWEEN ? AND ?", input.MinAge, input.MaxAge).
		Find(&users)

	if query.Error != nil {
		return nil, 0, query.Error
	}

	count := query.RowsAffected
	return users, count, nil
}
