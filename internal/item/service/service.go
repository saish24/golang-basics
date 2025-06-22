package service

import (
	"context"
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

func (i *ItemServiceImpl) GetItem(ctx context.Context, id string) (*models.Item, error) {
	if len(id) == 0 {
		return nil, errors.New("id is required")
	}

	itemObj, err := i.repository.GetItem(ctx, id)
	if err != nil {
		return nil, err
	}

	return itemObj.ToItem(), nil
}

func (i *ItemServiceImpl) AddItem(ctx context.Context, item *models.Item) error {
	return nil
}

func (i *ItemServiceImpl) UpdateItem(ctx context.Context, item *models.Item) error {
	return nil
}

func (i *ItemServiceImpl) DeleteItem(ctx context.Context, id string) error {
	return nil
}
