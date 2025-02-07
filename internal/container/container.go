package container

import (
	"context"

	"github.com/Igorsant/go-api/internal/container/components"
	"github.com/Igorsant/go-api/internal/database"
	"github.com/Igorsant/go-api/internal/domains/item"
)

type Container struct {
	Config      components.Config
	Provider    database.Provider
	ItemService item.Service
}

func NewContainer() Container {
	ctx := context.Background()
	cfg := components.NewConfig()
	provider := database.NewClient(ctx, cfg)
	itemRepo := item.NewRepository(&provider)
	itemService := item.NewService(itemRepo)

	return Container{
		Config:      cfg,
		Provider:    &provider,
		ItemService: itemService,
	}
}
