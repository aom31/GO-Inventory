package service

import (
	"context"

	"github.com/pkg/errors"

	"github.com/aom31/GO-Inventory/src/models"
	"github.com/aom31/GO-Inventory/src/repository"
)

type IUserService interface {
	FindUserById(userId string) (*models.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService() IUserService {
	return &userService{
		userRepository: repository.UserRepository{},
	}
}

func (service *userService) FindUserById(userId string) (*models.User, error) {
	user, err := service.userRepository.FindOneUser(context.Background(), userId)
	if err != nil {
		return nil, errors.Wrap(err, "user not found with id")
	}
	return user, nil
}
