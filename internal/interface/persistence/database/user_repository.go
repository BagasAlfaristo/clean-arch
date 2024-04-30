// user_repository.go

package database

import (
	"simple-project/internal/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) GetAllUsers() ([]domain.User, error) {
	var users []domain.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) AddUser(newUser domain.User) error {
	return r.DB.Create(&newUser).Error
}
