package item

import "context"

type Service interface {
	Get(ctx context.Context, id string) (Item, error)
	Update(ctx context.Context, id string) (Item, error)
	Delete(ctx context.Context, id string) (Item, error)
	Add(ctx context.Context, s Item) (Item, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

// Add implements Service.
func (s *service) Add(ctx context.Context, i Item) (Item, error) {
	return s.repo.Add(ctx, i)
}

// Delete implements Service.
func (s *service) Delete(ctx context.Context, id string) (Item, error) {
	panic("unimplemented")
}

// Get implements Service.
func (s *service) Get(ctx context.Context, id string) (Item, error) {
	return s.repo.Get(ctx, id)
}

// Update implements Service.
func (s *service) Update(ctx context.Context, id string) (Item, error) {
	panic("unimplemented")
}
