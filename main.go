package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"time"
)

type Transaction struct {
	ID                  int    `json:"id"`
	Amount              int    `json:"amount"`
	MessageType         string `json:"conversation_type"`
	CreatedAt           string `json:"created_at"`
	TransactionID       int    `json:"transaction_id"`
	PAN                 string `json:"pan"`
	TransactionCategory string `json:"transaction_category"`
	PostedTimestamp     string `json:"posted_timestamp"`
	TransactionType     string `json:"transaction_type"`
	SendingAccount      int    `json:"sending_account"`
	ReceivingAccount    int    `json:"receiving_account"`
	TransactionNote     string `json:"transaction_note"`
}

var transactions []Transaction

func maskPAN(pan string) string {
	if len(pan) <= 4 {
		return pan
	}

	masked := make([]rune, len(pan))
	for i, r := range pan {
		if len(pan)-i <= 4 {
			masked[i] = r
			continue
		}
		masked[i] = '*'
	}

	return string(masked)
}

func maskedTransactions(source []Transaction) []Transaction {
	masked := make([]Transaction, len(source))
	for i, t := range source {
		masked[i] = t
		masked[i].PAN = maskPAN(t.PAN)
	}
	return masked
}

func loadTransactions(path string) ([]Transaction, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading transactions file: %w", err)
	}

	var loaded []Transaction
	if err := json.Unmarshal(data, &loaded); err != nil {
		return nil, fmt.Errorf("decoding transactions: %w", err)
	}

	return loaded, nil
}

func respondJSON(w http.ResponseWriter, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetTransactions(w http.ResponseWriter, _ *http.Request) {
	respondJSON(w, maskedTransactions(transactions))
}

func parsePostedTimestamp(ts string) time.Time {
	parsed, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		return time.Time{}
	}
	return parsed
}

func GetTransactionsDescending(w http.ResponseWriter, _ *http.Request) {
	sorted := make([]Transaction, len(transactions))
	copy(sorted, transactions)

	sort.Slice(sorted, func(i, j int) bool {
		return parsePostedTimestamp(sorted[i].PostedTimestamp).After(parsePostedTimestamp(sorted[j].PostedTimestamp))
	})

	respondJSON(w, maskedTransactions(sorted))
}

func main() {
	transactionsPath := flag.String("transactions", "transactions.json", "path to JSON file containing transactions")
	flag.Parse()

	loaded, err := loadTransactions(*transactionsPath)
	if err != nil {
		log.Fatalf("failed to load transactions: %v", err)
	}
	transactions = loaded

	mux := http.NewServeMux()
	mux.HandleFunc("/transactions", GetTransactions)
	mux.HandleFunc("/transactions/posted-desc", GetTransactionsDescending)

	fmt.Println("Serving transactions on port 8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
