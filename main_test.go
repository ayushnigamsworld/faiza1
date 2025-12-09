package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
