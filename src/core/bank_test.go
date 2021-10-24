package bank

import "testing"

func TestAccount(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Micky",
			Address: "Greece",
			Phone:   "(30) 111 111",
		},
		Number:  2222,
		Balance: 100,
	}
	if account.Name == "" {
		t.Error("Can't create an Account object")
	}
}
