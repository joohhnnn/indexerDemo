package ethereum

import (
	"context"
	"database/sql"
	localABI"indexerDemo/indexer/abi"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"strings"
)


// ERC20TransferEventSignature is the signature of the Transfer event in an ERC20 contract
var ERC20TransferEventSignature = crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))

// TransferEvent represents an ERC20 Transfer event
type TransferEvent struct {
	From  common.Address
	To    common.Address
	Value *big.Int
}

type ERC20TransferEvent struct {
	BlockNumber int64
	TxHash      string
	FromAddress string
	ToAddress   string
	Value       string
}

// IndexPastEvents indexes past ERC20 transfer events from the Ethereum blockchain
func IndexPastEvents(client EthClientInterface, db *sql.DB, startBlock uint64, endBlock uint64) error {
	query := ethereum.FilterQuery{
		Topics: [][]common.Hash{
			{ERC20TransferEventSignature},
		},
		FromBlock: big.NewInt(int64(startBlock)),
		ToBlock:   big.NewInt(int64(endBlock)),
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return err
	}

	for _, vLog := range logs {
		transferEvent := TransferEvent{}
		err := parseLogIntoTransferEvent(vLog, &transferEvent)
		if err != nil {
			log.Println("Failed to parse log:", err)
			continue
		}

		// Store the event in the database
		err = StoreTransferEvent(db, int64(vLog.BlockNumber), vLog.TxHash.Hex(), &transferEvent)
		if err != nil {
			log.Println("Failed to store event:", err)
			continue
		}
	}

	return nil
}

// SubscribeToEvents subscribes to new ERC20 transfer events
func SubscribeToEvents(client EthClientInterface, db *sql.DB) (chan types.Log, error) {
	query := ethereum.FilterQuery{
		Topics: [][]common.Hash{
			{ERC20TransferEventSignature},
		},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err) // TODO: Handle error properly
			case vLog := <-logs:
				transferEvent := TransferEvent{}
				err := parseLogIntoTransferEvent(vLog, &transferEvent)
				if err != nil {
					log.Println("Failed to parse log:", err)
					continue
				}

				// Store the transfer event in the database
				err = StoreTransferEvent(db, int64(vLog.BlockNumber), vLog.TxHash.Hex(), &transferEvent)
				if err != nil {
					log.Println("Failed to store event:", err)
					continue
				}
			}
		}
	}()

	return logs, nil
}

// Parse log into TransferEvent
func parseLogIntoTransferEvent(vLog types.Log, event *TransferEvent) error {
	contractAbi, err := abi.JSON(strings.NewReader(localABI.ERC20TransferEventABI))
	if err != nil {
		return err
	}

	return contractAbi.UnpackIntoInterface(event, "Transfer", vLog.Data)
}

// Store the transfer event in the database
func StoreTransferEvent(db *sql.DB, blockNumber int64, txHash string, event *TransferEvent) error {
	// Convert to the database model
	dbEvent := ERC20TransferEvent{
		BlockNumber: blockNumber,
		TxHash:      txHash,
		FromAddress: event.From.String(),
		ToAddress:   event.To.String(),
		Value:       event.Value.String(),
	}

	query := `
	INSERT INTO erc20_transfers (block_number, tx_hash, from_address, to_address, value)
	VALUES (?, ?, ?, ?, ?)
	`
	_, err := db.Exec(query, dbEvent.BlockNumber, dbEvent.TxHash, dbEvent.FromAddress, dbEvent.ToAddress, dbEvent.Value)
	return err
}
