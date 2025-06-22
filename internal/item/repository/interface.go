package repository

import (
	"basics/internal/item/models"
	db "basics/pkg/dynamodb"
	"context"
)

type ItemRepository interface {
	GetItem(id string) (*models.DBItem, error)
	AddItem(item *models.DBItem) error
	UpdateItem(item *models.DBItem) error
	DeleteItem(id string) error
}

type ItemRepositoryImpl struct {
	dbClient *db.DatabaseClient
}

func NewItemRepository() *ItemRepositoryImpl {
	return &ItemRepositoryImpl{dbClient: db.GetDynamoClient(context.Background())}
}

func (i *ItemRepositoryImpl) GetItem(id string) (*models.DBItem, error) {
	//item, err := i.dbClient.GetItem(context.Background(), &db.GetItemInput{})
	return nil, nil
}

func (i *ItemRepositoryImpl) AddItem(item *models.DBItem) error {
	return nil
}

func (i *ItemRepositoryImpl) UpdateItem(item *models.DBItem) error {
	return nil
}

func (i *ItemRepositoryImpl) DeleteItem(id string) error {
	return nil
}
