package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

type Account struct {
	Username string `json:"username"`
	Email    string `json:"email`
}

func handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	var account Account
	slog.Info("Start handleCreateAccount()")
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		slog.Error("failed to decode the response body", "err", err)
		return
	}
	if err := notifyAccountCreated(account); err != nil {
		slog.Error("failed to notify account created", "err", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func notifyAccountCreated(account Account) error {
	slog.Info("Start notifyAccountCreated()")
	time.Sleep(time.Microsecond * 500)
	slog.Info("new account created", "username", account.Username, "email", account.Email)
	return nil
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/account", handleCreateAccount)

	slog.Info("Starting HTTP server on: 8082")

	err := http.ListenAndServe(":8082", mux)
	if err != nil {
		slog.Error("Error starting server: ", err)
	}
}
