package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/KNaiskes/bank"
)

var accounts = map[float64]*bank.Account{}

func main() {
	accounts[1001] = &bank.Account{
		Customer: bank.Customer{
			Name:    "Micky",
			Address: "Greece",
			Phone:   "111 222",
		},
		Number: 1001,
	}

	http.HandleFunc("/statement", statement)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func statement(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found", number)
		} else {
			fmt.Fprintf(w, account.Statement())
		}
	}
}
