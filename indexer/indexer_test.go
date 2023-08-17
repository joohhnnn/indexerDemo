package indexer

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/mock"
)

type MockEthClient struct {
	mock.Mock
}

func (m *MockEthClient) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	args := m.Called(ctx, q)
	return args.Get(0).([]types.Log), args.Error(1)
}

func (m *MockEthClient) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	args := m.Called(ctx, q, ch)
	return args.Get(0).(ethereum.Subscription), args.Error(1)
}
type MockSubscription struct {
	mock.Mock
}

func (m *MockSubscription) Err() <-chan error {
	return make(chan error)
}

func (m *MockSubscription) Unsubscribe() {}

func TestStartSubscribingToEvents(t *testing.T) {
	// Create a mock EthClient
	mockEthClient := new(MockEthClient)
	// Create a mock Subscription
	mockSubscription := new(MockSubscription)
	// Define the behavior of the mock EthClient
	mockEthClient.On("SubscribeFilterLogs", mock.Anything, mock.Anything, mock.Anything).Return(mockSubscription, nil)

	// Create a new Indexer with the mock EthClient
	indexer := NewIndexer(&Connections{EthClient: mockEthClient})

	// Call the startSubscribingToEvents method
	_, err := indexer.startSubscribingToEvents()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Assert that the expected methods were called on the mock EthClient
	mockEthClient.AssertExpectations(t)
}
