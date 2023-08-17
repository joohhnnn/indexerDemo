package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestHandleEvents(t *testing.T) {
	// Create a mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Prepare 60 mock data rows
	rows1 := sqlmock.NewRows([]string{"block_number", "tx_hash", "from_address", "to_address", "value"})
	for i := 0; i < 60; i++ {
		rows1.AddRow(1234+i, common.BytesToHash([]byte{byte(i)}), common.HexToAddress("0xAddress"), common.HexToAddress("0xAddress"), "100")
	}

	// Expected first database query
	mock.ExpectQuery("^SELECT (.+) FROM erc20_transfers WHERE").WillReturnRows(rows1)

	// Call the handleEvents function
	req, _ := http.NewRequest("GET", "/events?address=some_address", nil)
	rr := httptest.NewRecorder()
	handleEvents(db, rr, req)

	// Check the first HTTP response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Prepare 70 mock data rows
	rows2 := sqlmock.NewRows([]string{"block_number", "tx_hash", "from_address", "to_address", "value"})
	for i := 0; i < 70; i++ {
		rows2.AddRow(1234+i, common.BytesToHash([]byte{byte(i)}), common.HexToAddress("0xAddress"), common.HexToAddress("0xAddress"), "100")
	}

	// Expected second database query
	mock.ExpectQuery("^SELECT (.+) FROM erc20_transfers WHERE").WillReturnRows(rows2)

	// Call the handleEvents function again
	req, _ = http.NewRequest("GET", "/events?address=some_address", nil)
	rr = httptest.NewRecorder()
	handleEvents(db, rr, req)

	// Check the second HTTP response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Ensure all expected database calls have occurred
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
