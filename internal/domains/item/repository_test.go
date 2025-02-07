package item_test

import (
	"context"
	"testing"

	"github.com/Igorsant/go-api/internal/database"
	"github.com/Igorsant/go-api/internal/domains/item"
	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	repo := item.NewRepository(&database.MongoProvider{})
	i, err := repo.Update(context.TODO(), item.Item{})
	assert.NoError(t, err)
	assert.Equal(t, item.Item{}, i)
}
