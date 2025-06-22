package repository

import (
	"context"
	"golang-basics/internal/item/models"
	db "golang-basics/pkg/dynamodb"
)

type ItemRepositoryImpl struct {
	dbClient *db.DatabaseClient
}

func NewItemRepository() ItemRepository {
	return &ItemRepositoryImpl{dbClient: db.GetDynamoClient(context.Background())}
}

func (i *ItemRepositoryImpl) GetItem(ctx context.Context,id string) (*models.DBItem, error) {
	//item, err := i.dbClient.GetItem(ctx context.Context,context.Background(), &db.GetItemInput{})
	return nil, nil
}

func (i *ItemRepositoryImpl) AddItem(ctx context.Context,item *models.DBItem) error {
	return nil
}

func (i *ItemRepositoryImpl) UpdateItem(ctx context.Context,item *models.DBItem) error {
	return nil
}

func (i *ItemRepositoryImpl) DeleteItem(ctx context.Context,id string) error {
	return nil
}
