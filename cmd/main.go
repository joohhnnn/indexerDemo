package main

import (
	"indexerDemo/config"
	"indexerDemo/indexer"
	"log"
)

func main() {
	// Load the configuration
	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Establish connections (Ethereum client and database)
	connections, err := indexer.Connect(appConfig)
	if err != nil {
		log.Fatalf("Failed to establish connections: %v", err)
	}

	// Create the indexer
	idx := indexer.NewIndexer(connections)

	// Start the indexer
	idx.Start(appConfig)
}
