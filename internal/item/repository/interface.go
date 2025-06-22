package repository

import (
	"context"
	"golang-basics/internal/item/models"
)

type ItemRepository interface {
	GetItem(ctx context.Context, id string) (*models.DBItem, error)
	AddItem(ctx context.Context, item *models.DBItem) error
	UpdateItem(ctx context.Context, item *models.DBItem) error
	DeleteItem(ctx context.Context, id string) error
}
