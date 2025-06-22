package router

import (
	"basics/controller/item"
	"github.com/gin-gonic/gin"
)

func InitialiseGinServer() *gin.Engine {
	r := gin.New()
	initialiseRoutes(r)
	return r
}

func initialiseRoutes(r *gin.Engine) {
	itemController := item.NewItemController()

	r.GET("/item", itemController.GetItem)
	r.PUT("/item", itemController.InsertItem)
	r.POST("/item", itemController.UpdateItem)
	r.DELETE("/item", itemController.DeleteItem)
}
