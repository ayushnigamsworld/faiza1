package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestMaskPAN(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"1234567890123456", "************3456"},
		{"1234", "1234"},
		{"", ""},
	}

	for _, c := range cases {
		if got := maskPAN(c.input); got != c.expected {
			t.Fatalf("maskPAN(%s) = %s, expected %s", c.input, got, c.expected)
		}
	}
}

func TestGetTransactionsMasksPAN(t *testing.T) {
	transactions = []Transaction{
		{ID: 1, PAN: "1234567890123456", PostedTimestamp: "2020-06-11T19:11:24+00:00"},
	}

	req := httptest.NewRequest(http.MethodGet, "/transactions", nil)
	rr := httptest.NewRecorder()

	GetTransactions(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	var resp []Transaction
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(resp) != 1 {
		t.Fatalf("expected 1 transaction, got %d", len(resp))
	}

	if resp[0].PAN != "************3456" {
		t.Fatalf("expected masked PAN, got %s", resp[0].PAN)
	}
}

func TestGetTransactionsDescendingOrdersByPostedTimestamp(t *testing.T) {
	transactions = []Transaction{
		{ID: 1, PAN: "1111", PostedTimestamp: "2020-06-11T19:11:24+00:00"},
		{ID: 2, PAN: "2222", PostedTimestamp: "2020-06-11T19:11:25+00:00"},
		{ID: 3, PAN: "3333", PostedTimestamp: "2020-06-11T19:11:23+00:00"},
	}

	req := httptest.NewRequest(http.MethodGet, "/transactions/posted-desc", nil)
	rr := httptest.NewRecorder()

	GetTransactionsDescending(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	var resp []Transaction
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	expectedOrder := []int{2, 1, 3}
	for i, id := range expectedOrder {
		if resp[i].ID != id {
			t.Fatalf("expected transaction %d at position %d, got %d", id, i, resp[i].ID)
		}
	}
}

func TestLoadTransactionsFromCustomFile(t *testing.T) {
	tmp := t.TempDir()

	data := `[
{"id":10,"amount":5000,"conversation_type":"payment","created_at":"2020-06-11T19:11:22+00:00","transaction_id":123,"pan":"4567123412341234","transaction_category":"groceries","posted_timestamp":"2020-06-11T19:11:22+00:00","transaction_type":"debit","sending_account":123,"receiving_account":456,"transaction_note":"temporary"}
]`

	path := tmp + "/custom_transactions.json"
	if err := os.WriteFile(path, []byte(data), 0o644); err != nil {
		t.Fatalf("failed to write temp transactions: %v", err)
	}

	loaded, err := loadTransactions(path)
	if err != nil {
		t.Fatalf("expected to load transactions, got error: %v", err)
	}

	if len(loaded) != 1 || loaded[0].ID != 10 {
		t.Fatalf("expected 1 transaction with ID 10, got %+v", loaded)
	}

	if loaded[0].PAN != "4567123412341234" {
		t.Fatalf("expected PAN to be read from file, got %s", loaded[0].PAN)
	}
}

func TestLoadTransactionsMissingFile(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "missing.json")

	if _, err := loadTransactions(missing); !errors.Is(err, ErrTransactionsFileMissing) {
		t.Fatalf("expected ErrTransactionsFileMissing, got %v", err)
	}
}

func TestLoadTransactionsMissingField(t *testing.T) {
	tmp := t.TempDir()

	data := `[{
"id":1,
"amount":0,
"conversation_type":"payment",
"created_at":"2020-06-11T19:11:22+00:00",
"transaction_id":123,
"pan":"4567123412341234",
"transaction_category":"groceries",
"posted_timestamp":"2020-06-11T19:11:22+00:00",
"transaction_type":"debit",
"sending_account":123,
"receiving_account":456,
"transaction_note":"temporary"}]`

	path := filepath.Join(tmp, "invalid_transactions.json")
	if err := os.WriteFile(path, []byte(data), 0o644); err != nil {
		t.Fatalf("failed to write temp transactions: %v", err)
	}

	if _, err := loadTransactions(path); !errors.Is(err, ErrTransactionMissingField) {
		t.Fatalf("expected ErrTransactionMissingField, got %v", err)
	}
}
