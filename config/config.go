package config

import (
	"flag"
	"os"
)

// AppConfig holds the configuration for the application
type AppConfig struct {
	RPCURL         string
	IndexPast      bool
	Subscribe      bool
	DatabaseDSN    string
	APIPort        string
	StartBlock     uint64
	BlocksPerCycle uint64
}

// LoadConfig parses the command-line flags and returns the application configuration
func LoadConfig() (*AppConfig, error) {
	rpcURL 			:= flag.String("rpcurl", "http://localhost:8545", "Ethereum RPC URL")
	indexPast 		:= flag.Bool("indexpast", false, "Index past ERC20 transfer events")
	subscribe 		:= flag.Bool("subscribe", false, "Subscribe to new ERC20 transfer events")
	databaseDSN 	:= flag.String("db", "sqlite3.db", "Database connection string")
	apiPort 		:= flag.String("apiport", "8080", "API server port")
	startBlock 		:= flag.Uint64("startblock", 0, "Starting block for indexing past events")
	blocksPerCycle	:= flag.Uint64("blockspencycle", 10, "Number of blocks per cycle when indexing past events")
	flag.Parse()

	// Validate the flags
	if !*indexPast && !*subscribe {
		return nil, os.ErrInvalid
	}

	return &AppConfig{
		RPCURL:         *rpcURL,
		IndexPast:      *indexPast,
		Subscribe:      *subscribe,
		DatabaseDSN:    *databaseDSN,
		APIPort:        *apiPort,
		StartBlock:     *startBlock,
		BlocksPerCycle: *blocksPerCycle,
	}, nil
}
