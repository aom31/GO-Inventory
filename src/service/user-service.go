package service

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/aom31/GO-Inventory/src/models"
	"github.com/aom31/GO-Inventory/src/repository"
)

type IUserService interface {
	FindUserById(ctx context.Context, userId string) (*models.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(Client *mongo.Client) IUserService {
	return &userService{
		userRepository: *repository.NewUserRepository(Client),
	}
}

func (service *userService) FindUserById(ctx context.Context, userId string) (*models.User, error) {
	user, err := service.userRepository.FindOneUser(context.Background(), userId)
	if err != nil {
		return nil, errors.Wrap(err, "user not found with id")
	}
	return user, nil
}
