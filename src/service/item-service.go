package service

import "github.com/aom31/GO-Inventory/src/repository"

type IItemService interface {
}

type itemService struct {
	itemRepository repository.ItemRepository
}

func NewItemService() IItemService {
	return &itemService{
		itemRepository: repository.ItemRepository{},
	}
}
