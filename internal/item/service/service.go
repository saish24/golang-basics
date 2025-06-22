package service

import (
	"errors"
	"golang-basics/internal/item/models"
	"golang-basics/internal/item/repository"
)

type ItemServiceImpl struct {
	repository repository.ItemRepository
}

func NewItemService(repository repository.ItemRepository) ItemService {
	return &ItemServiceImpl{
		repository: repository,
	}
}

func (i *ItemServiceImpl) GetItem(id string) (*models.Item, error) {
	if len(id) == 0 {
		return nil, errors.New("id is required")
	}

	itemObj, err := i.repository.GetItem(id)
	if err != nil {
		return nil, err
	}

	return itemObj.ToItem(), nil
}

func (i *ItemServiceImpl) AddItem(item *models.Item) error {
	return nil
}

func (i *ItemServiceImpl) UpdateItem(item *models.Item) error {
	return nil
}

func (i *ItemServiceImpl) DeleteItem(id string) error {
	return nil
}
