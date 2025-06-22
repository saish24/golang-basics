package system

import (
	"context"
	"golang-basics/router"

	"github.com/gin-gonic/gin"
)

func InitialiseGinServer(ctx context.Context) (*gin.Engine, error) {
	server := router.InitialiseGinServer()
	return server, nil
}
