package database

import (
	"context"
)

//go:generate mockgen -source contract.go -destination ../mocks/mock_provider.go -package=mocks github.com/Igorsant/go-api/internal/database *
type Provider interface {
	Insert(context.Context, string, interface{}) error
	Get(ctx context.Context, collectionName string, filter interface{}, result interface{}) error
	Disconnect(context.Context) error
}
