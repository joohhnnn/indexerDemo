package indexer

import (
	"database/sql"
	"indexerDemo/config"
	"indexerDemo/indexer/database"
	"indexerDemo/indexer/ethereum"
)

type Connections struct {
	EthClient      ethereum.EthClientInterface 
	DB             *sql.DB
}

func Connect(appConfig *config.AppConfig) (*Connections, error) {
	client, err := ethereum.ConnectEthereum(appConfig.RPCURL)
	if err != nil {
		return nil, err
	}

	db, err := database.ConnectDatabase(appConfig.DatabaseDSN)
	if err != nil {
		return nil, err
	}

	if err := database.InitSchema(db); err != nil {
		return nil, err
	}

	return &Connections{
		EthClient: client,
		DB:        db,
	}, nil
}
