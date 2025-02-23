package item_test

import (
	"context"
	"testing"

	"github.com/Igorsant/go-api/internal/domains/item"
	"github.com/Igorsant/go-api/internal/mocks"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/mock/gomock"
)

type ItemRepositoryTestSuite struct {
	suite.Suite
	repo         item.Repository
	providerMock *mocks.MockProvider
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(ItemRepositoryTestSuite))
}

func (s *ItemRepositoryTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(s.T())
	defer mockCtrl.Finish()

	providerMock := mocks.NewMockProvider(mockCtrl)
	s.repo = item.NewRepository(providerMock)
	s.providerMock = providerMock
}

func (s *ItemRepositoryTestSuite) TestGet() {
	s.providerMock.EXPECT().Get(context.TODO(), "items", bson.D{{Key: "id", Value: "id"}}, &item.Item{}).Return(nil)
	itemObj, err := s.repo.Get(context.TODO(), "id")
	s.NoError(err)
	s.Equal(item.Item{}, itemObj)
}
