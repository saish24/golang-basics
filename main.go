package main

import (
	"golang-basics/internal/app_init"
	"golang-basics/internal/system"
	"net/http"
)

func main() {
	ctx := app_init.Init()
	defer app_init.CloseConnections(ctx)

	handler, ginErr := system.InitialiseGinServer(ctx)
	if ginErr != nil {
		panic(ginErr)
	}

	server := http.Server{
		Addr:    ":3001",
		Handler: handler,
	}

	if serverErr := server.ListenAndServe(); serverErr != nil {
		panic(serverErr)
	}
}
