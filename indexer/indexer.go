package indexer

import (
	"indexerDemo/config"
	"indexerDemo/indexer/api"
	"indexerDemo/indexer/ethereum"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
)


type Indexer struct {
	Connections *Connections
}

func NewIndexer(connections *Connections) *Indexer {
	return &Indexer{Connections: connections}
}

func (i *Indexer) Start(appConfig *config.AppConfig) {
	var firstSubscribedBlock uint64

	if appConfig.Subscribe {
		logs, err := i.startSubscribingToEvents()
		if err != nil {
			log.Fatalf("Failed to subscribe to new events: %v", err)
		}

		// Wait for the first subscribed block or timeout
		select {
		case firstLog := <-logs:
			firstSubscribedBlock = firstLog.BlockNumber
		case <-time.After(5 * time.Minute): // Adjust the timeout as needed
			log.Fatalf("Timed out waiting for the first subscribed block")
		}
	}

	if appConfig.IndexPast {
		if err := i.startIndexingPastEvents(appConfig.StartBlock, firstSubscribedBlock-1, appConfig.BlocksPerCycle); err != nil {
			log.Fatalf("Failed to index past events: %v", err)
		}
	}

	i.startAPI(appConfig.APIPort)
}

func (i *Indexer) startIndexingPastEvents(startBlock uint64, firstSubscribedBlock uint64, blocksPerCycle uint64) error {
	for block := firstSubscribedBlock; block >= startBlock; block -= blocksPerCycle {
		endBlock := block
		startBlockCycle := endBlock - blocksPerCycle + 1
		if startBlockCycle < startBlock {
			startBlockCycle = startBlock
		}

		err := ethereum.IndexPastEvents(i.Connections.EthClient, i.Connections.DB, startBlockCycle, endBlock)
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *Indexer) startSubscribingToEvents() (chan types.Log, error) {
	return ethereum.SubscribeToEvents(i.Connections.EthClient, i.Connections.DB)
}

func (i *Indexer) startAPI(port string) {
	api.StartServer(i.Connections.DB, port)
}
