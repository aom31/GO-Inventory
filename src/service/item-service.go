package service

import (
	"context"

	"github.com/aom31/GO-Inventory/src/models"
	"github.com/aom31/GO-Inventory/src/repository"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type IItemService interface {
	FindItems(ctx context.Context, items *[]models.Item) error
	FindOneItem(c context.Context, itemId string) (*models.Item, error)
}

type itemService struct {
	itemRepository repository.ItemRepository
}

func NewItemService(Client *mongo.Client) IItemService {
	return &itemService{
		itemRepository: *repository.NewItemRepository(Client),
	}
}

func (service *itemService) FindItems(c context.Context, items *[]models.Item) error {
	if err := service.itemRepository.FindItems(c, items); err != nil {
		return errors.Wrap(err, "not found item")
	}

	return nil
}

func (service *itemService) FindOneItem(c context.Context, itemId string) (*models.Item, error) {
	item, err := service.itemRepository.FindOneItem(c, itemId)
	if err != nil {
		return nil, errors.Wrapf(err, "not found item with id :%v", itemId)
	}

	return item, nil
}
