package service

import "golang-basics/internal/item/models"

type ItemService interface {
	GetItem(id string) (*models.Item, error)
	AddItem(item *models.Item) error
	UpdateItem(item *models.Item) error
	DeleteItem(id string) error
}
