package item

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	client mongo.Client
}

func NewRepository(mongoClient mongo.Client) Repository {
	return repository{
		client: mongoClient,
	}
}

const (
	databaseName   = "go-api"
	collectionName = "items"
)

func (r repository) Get(ctx context.Context, id string) (Item, error) {
	collection := r.client.Database(databaseName).Collection(collectionName)

	filter := bson.D{{Key: "id", Value: id}}
	var item Item

	singleResult := collection.FindOne(ctx, filter)
	if singleResult.Err() != nil {
		return Item{}, singleResult.Err()
	}

	singleResult.Decode(&item)
	return item, nil
}

func (r repository) Update(ctx context.Context, i Item) (Item, error)    { return Item{}, nil }
func (r repository) Delete(ctx context.Context, id string) (Item, error) { return Item{}, nil }

func (r repository) Add(ctx context.Context, i Item) (Item, error) {
	collection := r.client.Database(databaseName).Collection(collectionName)

	_, err := collection.InsertOne(ctx, i)
	if err != nil {
		return Item{}, nil
	}

	return i, nil
}
