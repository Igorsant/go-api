package item

import "context"

type Repository interface {
	Get(ctx context.Context, id string) (Item, error)
	Update(ctx context.Context, i Item) (Item, error)
	Delete(ctx context.Context, id string) (Item, error)
	Add(ctx context.Context, i Item) (Item, error)
}

type Item struct {
	ID    string
	Model string
	Qtd   int
}
