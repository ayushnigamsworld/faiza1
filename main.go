package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"
    )

type Transaction struct {
  ID             int `json:"id"`
  Amount         int `json:"amount"`
	MessageType    string `json:"conversation_type"`
	CreatedAt      string `json:"created_at"`
  TransactionID  int `json:"created_at"`
  PAN            int `json:"pan"`
  TransactionCategory string `json:"transaction_category"`
  PostedTimeStamp string `json:"posted_timestamp"`
  TransactionType string `json:"transaction_type"`
  SendingAccount  int   `json:"sending_account"`
  ReceivingAccount int `json:"receiving_account"`
  TransactionNote string `json:"transaction_note"`
}

var transactions []Transaction

func GetTransactions(w http.ResponseWriter, r *http.Request) {
  //mock data
  transactions = append(transactions, Transaction{ID: 1, Amount: 200, MessageType: "Debit", CreatedAt: "2020-06-11T19:11:24+00:00", TransactionID: 10, PAN: 4080230386144446, TransactionCategory: "Grocery", PostedTimeStamp: "2020-06-11T19:11:24+00:00", TransactionType: "POS", SendingAccount: 39203, ReceivingAccount: 993020, TransactionNote: "Merchant 00308281"}, Transaction{ID: 2, Amount: 499, MessageType: "Credit", CreatedAt: "2020-06-11T19:11:24+00:00", TransactionID: 12, PAN: 5166697943434128, TransactionCategory: "Food and Beverage", PostedTimeStamp: "2020-06-11T19:11:24+00:00", TransactionType: "POS", SendingAccount: 39400, ReceivingAccount: 9233020, TransactionNote: "Jimmys Corn and Cheese refund"}, Transaction{ID: 3, Amount: 20000, MessageType: "Debit", CreatedAt: "2020-06-11T19:11:24+00:00", TransactionID: 17, PAN: 5488452462266852, TransactionCategory: "ATM", PostedTimeStamp: "2020-06-11T19:11:24+00:00", TransactionType: "POS", SendingAccount: 99302, ReceivingAccount: 11209, TransactionNote: "ATM #39902 Burrard Street"}, Transaction{ID: 4, Amount: 8839, MessageType: "Debit", CreatedAt: "2020-06-11T19:11:24+00:00", TransactionID: 10, PAN: 4954335252282726, TransactionCategory: "Automotive", PostedTimeStamp: "2020-06-11T19:11:24+00:00", TransactionType: "POS", SendingAccount: 83839, ReceivingAccount: 9233020, TransactionNote: "Muffler Bearings Inc."}, Transaction{ID: 5, Amount: 6173, MessageType: "Debit", CreatedAt: "2020-06-21T20:11:24+00:00", TransactionID: 21, PAN: 4844085301308048, TransactionCategory: "Household", PostedTimeStamp: "2020-06-21T20:11:24+00:00", TransactionType: "POS", SendingAccount: 9018, ReceivingAccount: 9222020, TransactionNote: "Amazon Inc."}, Transaction{ID: 6, Amount: 9018, MessageType: "Debit", CreatedAt: "2020-06-14T21:11:24+00:00", TransactionID: 33, PAN: 4090070794938361, TransactionCategory: "Electronics", PostedTimeStamp: "2020-06-14T21:11:24+00:00", TransactionType: "POS", SendingAccount: 83339, ReceivingAccount: 9233021, TransactionNote: "Apple Store"}, Transaction{ID: 7, Amount: 1275, MessageType: "Debit", CreatedAt: "2020-06-09T11:11:24+00:00", TransactionID: 12, PAN: 4807678678904632, TransactionCategory: "Cryptocurrency", PostedTimeStamp: "2020-06-09T11:11:24+00:00", TransactionType: "POS", SendingAccount: 83839, ReceivingAccount: 9233020, TransactionNote: "Bittrex Inc."}, Transaction{ID: 8, Amount: 2167, MessageType: "Debit", CreatedAt: "2020-06-18T15:11:24+00:00", TransactionID: 44, PAN: 4673062314928753, TransactionCategory: "Food and Beverage", PostedTimeStamp: "2020-06-18T15:11:24+00:00", TransactionType: "POS", SendingAccount: 8569, ReceivingAccount: 533020, TransactionNote: "Giorgios Pizza Ltd."}, Transaction{ID: 9, Amount: 3178, MessageType: "Debit", CreatedAt: "2020-05-01T14:11:24+00:00", TransactionID: 90, PAN: 5109473381765575, TransactionCategory: "Internet Services", PostedTimeStamp: "2020-05-01T14:11:24+00:00", TransactionType: "POS", SendingAccount: 63639, ReceivingAccount: 4233010, TransactionNote: "Google"}, Transaction{ID: 10, Amount: 7718, MessageType: "Debit", CreatedAt: "2020-07-11T19:11:24+00:00", TransactionID: 109, PAN: 5158563621617519, TransactionCategory: "Health Services", PostedTimeStamp: "2020-07-11T19:11:25+00:00", TransactionType: "POS", SendingAccount: 13839, ReceivingAccount: 244020, TransactionNote: "Vons Nails"})


  fmt.Print("Current Transactions:\n")
  json.NewEncoder(w).Encode(transactions)
  
}



func main() {
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/transactions", GetTransactions).Methods("GET")
  fmt.Printf("Serving transactions on port 8000")
  log.Fatal(http.ListenAndServe(":8000", router))
}