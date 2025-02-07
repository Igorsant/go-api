package item

import (
	"context"

	"github.com/Igorsant/go-api/internal/database"
	"go.mongodb.org/mongo-driver/bson"
)

type repository struct {
	provider database.Provider
}

func NewRepository(provider database.Provider) Repository {
	return repository{
		provider: provider,
	}
}

const (
	collectionName = "items"
)

func (r repository) Get(ctx context.Context, id string) (Item, error) {
	filter := bson.D{{Key: "id", Value: id}}
	var item Item

	err := r.provider.Get(ctx, collectionName, filter, &item)
	if err != nil {
		return Item{}, err
	}

	return item, nil
}

func (r repository) Update(ctx context.Context, i Item) (Item, error)    { return Item{}, nil }
func (r repository) Delete(ctx context.Context, id string) (Item, error) { return Item{}, nil }

func (r repository) Add(ctx context.Context, i Item) (Item, error) {

	err := r.provider.Insert(ctx, collectionName, i)
	if err != nil {
		return Item{}, err
	}

	return i, nil
}
