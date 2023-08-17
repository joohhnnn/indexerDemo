package database

import (
	"testing"
	"reflect"
)

func TestDatabaseOperations(t *testing.T) {
	// Connect to an in-memory database
	db, err := ConnectDatabase(":memory:")
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Initialize the schema
	err = InitSchema(db)
	if err != nil {
		t.Fatalf("Failed to initialize schema: %v", err)
	}

	// Create an ERC20 transfer event for testing
	testEvent := &ERC20TransferEvent{
		BlockNumber:  12345,
		TxHash:       "testHash",
		FromAddress:  "fromAddress",
		ToAddress:    "toAddress",
		Value:        "1000",
	}

	// Store the event
	err = StoreEvent(db, testEvent)
	if err != nil {
		t.Fatalf("Failed to store event: %v", err)
	}

	// Retrieve the event
	events, err := RetrieveEvents(db, "fromAddress", 1)
	if err != nil {
		t.Fatalf("Failed to retrieve events: %v", err)
	}

	// Check if the retrieved event matches the stored event
	if !reflect.DeepEqual(events[0], testEvent) {
		t.Fatalf("Mismatched events: got %+v, want %+v", events[0], testEvent)
	}
}
