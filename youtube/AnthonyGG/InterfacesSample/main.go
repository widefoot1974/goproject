package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

type AccountNotifier interface {
	NotifyAccountCreated(context.Context, Account) error
}

type SimpleAccountNotifier struct{}

func (n SimpleAccountNotifier) NotifyAccountCreated(ctx context.Context, accout Account) error {
	slog.Info("new account created", "username", accout.Username, "email", accout.Email)
	return nil
}

type AccountHandler struct {
	AccountNotifier AccountNotifier
}

type Account struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (h *AccountHandler) handleCreateAccount(w http.ResponseWriter, r *http.Request) {

	slog.Info("Start handleCreateAccount()", "Method", r.Method)

	var account Account
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		slog.Error("failed to decode the response body", "err", err)
		return
	}
	// Logic
	if err := h.AccountNotifier.NotifyAccountCreated(r.Context(), account); err != nil {
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

	accountHandler := &AccountHandler{
		AccountNotifier: SimpleAccountNotifier{},
	}

	mux.HandleFunc("/account", accountHandler.handleCreateAccount)

	slog.Info("Starting HTTP server on: 8082")

	err := http.ListenAndServe(":8082", mux)
	if err != nil {
		slog.Error("Error starting server: ", err)
	}
}
