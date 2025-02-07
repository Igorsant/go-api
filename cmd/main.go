package main

import (
	"context"

	"github.com/Igorsant/go-api/internal/container"
	"github.com/Igorsant/go-api/internal/domains/item"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := gin.Default()

	container := container.NewContainer()
	defer func() {
		if err := container.Provider.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	item.RegisterRoutes(container.ItemService, r)
	r.Run(container.Config.Port)
}
