package main

import (
	"fmt"
	"net/http"

	httpHandler "uptime-monitor/internal/http"
	"uptime-monitor/internal/storage"
)

func main() {
	store := storage.NewMemoryStorage()
	handler := httpHandler.NewHandler(store)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/checks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateCheck(w, r)
		case http.MethodGet:
			handler.ListChecks(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Servidor rodando na porta 8080")
	http.ListenAndServe(":8080", nil)
}