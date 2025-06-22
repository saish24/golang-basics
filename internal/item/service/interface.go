package service

import (
	"context"
	"golang-basics/internal/item/models"
)

type ItemService interface {
	GetItem(ctx context.Context, id string) (*models.Item, error)
	AddItem(ctx context.Context, item *models.Item) error
	UpdateItem(ctx context.Context, item *models.Item) error
	DeleteItem(ctx context.Context, id string) error
}
