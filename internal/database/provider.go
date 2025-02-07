package database

import "context"

type Provider interface {
	Insert(context.Context, string, interface{}) error
	Get(ctx context.Context, collectionName string, filter interface{}, result interface{}) error
	Disconnect(context.Context) error
}
