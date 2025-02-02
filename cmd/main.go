package main

import (
	"context"
	"fmt"

	"github.com/Igorsant/go-api/internal/container/components"
	"github.com/Igorsant/go-api/internal/domains/item"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	godotenv.Load()
	r := gin.Default()
	cfg := components.NewConfig()

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb://root:example@localhost:27017").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	itemRepo := item.NewRepository(*client)
	itemService := item.NewService(itemRepo)
	item.RegisterRoutes(itemService, r)
	fmt.Println(cfg.Port)
	r.Run()
}
