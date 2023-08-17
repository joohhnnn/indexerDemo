package database

import (
	"database/sql"
)

// StoreEvent stores an ERC20 transfer event in the database
func StoreEvent(db *sql.DB, event *ERC20TransferEvent) error {
	query := `
	INSERT INTO erc20_transfers (block_number, tx_hash, from_address, to_address, value)
	VALUES (?, ?, ?, ?, ?)
	`
	_, err := db.Exec(query, event.BlockNumber, event.TxHash, event.FromAddress, event.ToAddress, event.Value)
	if err != nil {
		return err
	}
	return nil
}

// RetrieveEvents retrieves ERC20 transfer events from the database
func RetrieveEvents(db *sql.DB, address string, limit int) ([]*ERC20TransferEvent, error) {
	query := `
	SELECT block_number, tx_hash, from_address, to_address, value
	FROM erc20_transfers
	WHERE from_address = ?
	LIMIT ?
	`
	rows, err := db.Query(query, address, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*ERC20TransferEvent
	for rows.Next() {
		var event ERC20TransferEvent
		if err := rows.Scan(&event.BlockNumber, &event.TxHash, &event.FromAddress, &event.ToAddress, &event.Value); err != nil {
			return nil, err
		}
		events = append(events, &event)
	}
	return events, nil
}
