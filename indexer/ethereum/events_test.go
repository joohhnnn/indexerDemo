package ethereum_test

import (
	"context"
	"database/sql"
	localEthereum"indexerDemo/indexer/ethereum"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
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

func TestIndexPastEvents(t *testing.T) {
	// Create a mock Ethereum client
	mockEthClient := new(MockEthClient)

	// Define the expected logs
	expectedLogs := []types.Log{
		{
			Address: common.HexToAddress("0xAddress1"),
		},
		{
			Address: common.HexToAddress("0xAddress2"),
		},
	}

	// Expect FilterLogs to be called with specific arguments and return the expected logs
	mockEthClient.On("FilterLogs", mock.Anything, mock.Anything).Return(expectedLogs, nil)

	// Create a mock database connection (you can replace this with a real connection if needed)
	db, _ := sql.Open("sqlite3", ":memory:")

	// Call the function being tested
	err := localEthereum.IndexPastEvents(mockEthClient, db, 0, 100)
	if err != nil {
		t.Fatalf("IndexPastEvents failed: %v", err)
	}

	// Assert that the expected calls were made to the mock Ethereum client
	mockEthClient.AssertExpectations(t)

	// Additional assertions to verify the database state can be added here
}
