package database

import (
	"context"
	"fmt"

	"github.com/Igorsant/go-api/internal/container/components"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoProvider struct {
	client mongo.Client
	config components.Config
}

func NewClient(ctx context.Context, cfg components.Config) MongoProvider {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(cfg.MongoURI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	go func() {
		err := client.Ping(ctx, nil)
		if err != nil {
			panic(err)
		}
	}()
	return MongoProvider{
		client: *client,
		config: cfg,
	}
}

func (m MongoProvider) Disconnect(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}

func (m *MongoProvider) Get(ctx context.Context, collection string, filter interface{}, result interface{}) error {
	singleResult := m.client.Database(m.config.DatabaseName).Collection(collection).FindOne(ctx, filter)

	if singleResult.Err() != nil {
		return singleResult.Err()
	}
	singleResult.Decode(result)
	return nil
}

func (m *MongoProvider) Insert(ctx context.Context, collection string, item interface{}) error {
	_, err := m.client.Database(m.config.DatabaseName).Collection(collection).InsertOne(ctx, item)
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}
