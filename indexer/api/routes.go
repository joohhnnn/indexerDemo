package api

import (
	"encoding/json"
	"database/sql"
	"net/http"
	"indexerDemo/indexer/database"
)

// handleEvents handles the /events endpoint
func handleEvents(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Address parameter is required", http.StatusBadRequest)
		return
	}

	events, err := database.RetrieveEvents(db, address, 50)
	if err != nil {
		http.Error(w, "Failed to retrieve events", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(events)
}
