package api

import (
	"database/sql"
	"log"
	"net/http"
)

// StartServer starts the HTTP server
func StartServer(db *sql.DB, port string) {
	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		handleEvents(db, w, r) // 调用之前定义的handleEvents函数
	})

	log.Printf("Server started on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
