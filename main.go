package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	router := gin.Default()
	setupRoutes(router)
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}

func setupRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		//c.JSON(200, response{})

		request, err := http.NewRequest(http.MethodGet, "http://yahoo.com", nil)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		response, err := http.DefaultClient.Do(request)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		defer http.DefaultClient.CloseIdleConnections()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		c.JSON(http.StatusOK, string(body))
	})
}

type response struct {
	Message string `json:"message,omitempty"`
}
