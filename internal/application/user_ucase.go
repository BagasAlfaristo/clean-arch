package application

import (
	"simple-project/internal/domain"
	"simple-project/internal/interface/persistence/database"
)

type UserUsecase interface {
	GetAllUsers() ([]domain.User, error)
	AddUser(newUser domain.User) error
}

type userUsecase struct {
	userRepo database.UserRepository
}

func NewUserUsecase(repo database.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: repo,
	}
}

func (uc *userUsecase) GetAllUsers() ([]domain.User, error) {
	return uc.userRepo.GetAllUsers()
}

func (uc *userUsecase) AddUser(newUser domain.User) error {
	return uc.userRepo.AddUser(newUser)
}
