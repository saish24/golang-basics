package item

import (
	"golang-basics/internal/item/models"
	"golang-basics/internal/item/repository"
	"golang-basics/internal/item/service"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	ItemService service.ItemService
}

func NewItemController() *Controller {
	itemRepo := repository.NewItemRepository()
	itemService := service.NewItemService(itemRepo)

	return &Controller{ItemService: itemService}
}

func (ic *Controller) GetItem(c *gin.Context) {
	id := c.Query("id")
	item, err := ic.ItemService.GetItem(c, id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"item": item})
}

func (ic *Controller) InsertItem(c *gin.Context) {
	item := &models.Item{}
	if err := c.ShouldBind(item); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ic.ItemService.AddItem(c, item); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"item": item})
}

func (ic *Controller) UpdateItem( c *gin.Context) {
	item := &models.Item{}
	if err := c.ShouldBind(item); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ic.ItemService.UpdateItem(c, item); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"item": item})
}

func (ic *Controller) DeleteItem( c *gin.Context) {
	id := c.Query("id")
	if err := ic.ItemService.DeleteItem(c, id); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Item deleted successfully"})
}
