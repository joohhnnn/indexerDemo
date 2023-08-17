package ethereum

import (
	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthClientInterface interface {
	FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error)
	SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)
}

// ConnectEthereum establishes a connection to the Ethereum network
// connectEthereum connects to an Ethereum client at the given URL
func ConnectEthereum(url string) (EthClientInterface, error) {
	return ethclient.Dial(url)
}
